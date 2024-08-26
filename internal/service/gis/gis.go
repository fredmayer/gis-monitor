package gis

import (
	"context"
	"encoding/json"
	"fmt"
	"gis-crawler/internal/config"
	"gis-crawler/internal/models"
	"gis-crawler/internal/rest"
	"gis-crawler/pkg/logging"
)

type Service struct {
	store    StoreInterface
	client   ClientInterface
	settings config.Gis
}

func New(store StoreInterface, client ClientInterface, settings config.Gis) *Service {
	return &Service{
		store, client, settings,
	}
}

func (s *Service) Handle() {
	data := s.Parse()

	logging.Log().Infof("fetched %d lots \n", len(*data))
	if len(*data) == 0 {
		logging.Log().Warnln("parsed data from url is null")
		return
	}

	for _, row := range *data {
		insert, update := s.Check(&row)
		if insert {
			logging.Log().Infof("обнаружен новый лот: %s \n %s \n ====== \n", row.LotName, row.LotDescription)
			s.Save(&row, false)
			s.Notify(&row, false)
		}
		if update {
			logging.Log().Infof("изменился статус в лоте: %s на %s \n", row.LotName, row.LotStatus)
			s.Save(&row, true)
			s.Notify(&row, true)
		}

	}

	logging.Get().Debugln("done")
}

// Parse - парсим gis сайт по параметрам из конфига
func (s *Service) Parse() *[]models.LotDto {
	for key, value := range s.settings.Params {
		s.client.AddParam(key, value)
	}

	s.client.CreateRequest("GET", s.settings.Endpoint)
	data := make([]models.LotDto, 0)

	resp := s.getData(1)
	if len(resp.Content) > 0 {
		for _, row := range resp.Content {
			data = append(data, row)
		}
	}

	if resp.TotalPages > 1 {
		for i := 2; i <= resp.TotalPages; i++ {
			resp = s.getData(i)
			if len(resp.Content) > 0 {
				for _, row := range resp.Content {
					data = append(data, row)
				}
			}
		}
	}

	return &data
}

// getData отправка запроса на сайт gis с получением ответа
func (s *Service) getData(page int) *models.ContentOfLotDto {
	s.client.AddParam("page", string(page))
	var resp models.ContentOfLotDto
	body := s.client.Send()
	if body != nil {
		err := json.Unmarshal(body, &resp)
		if err != nil {
			logging.Get().Errorln("json response error: ", err, ": ") //, string(body[:])
		}
	}

	return &resp
}

// Check - проверяем были есть ли лоты в БД
func (s *Service) Check(dto *models.LotDto) (bool, bool) {
	lot, ok := s.store.GetByUID(dto.ID)
	if !ok {
		return true, false
	}

	if lot.Status != dto.LotStatus {
		return false, true
	}

	return false, false
}

// Save - сохраняем лоты
func (s *Service) Save(dto *models.LotDto, update bool) {
	model := models.Lot{
		ID:     dto.ID,
		Status: dto.LotStatus,
	}

	if update {
		m, _ := s.store.GetByUID(dto.ID)
		m.Status = dto.LotStatus
		err := s.store.Update(m)
		if err != nil {
			logging.Log().Errorln("update error: ", err)
		}
	} else {
		err := s.store.Save(&model)
		if err != nil {
			logging.Log().Errorln("save error: ", err)
		}
	}
}

// Notify - уведомляем о новых лотах
func (s *Service) Notify(lot *models.LotDto, update bool) {
	var template string

	if !update {
		template = fmt.Sprintf("<strong>🟢 Новый лот: %s</strong> \n", lot.LotName)
	} else {
		template = fmt.Sprintf("<strong>🌕 Лот обновлен: %s</strong> \n", lot.LotName)
	}

	link := s.settings.Host + s.settings.Link + lot.ID
	linkStr := fmt.Sprintf("\n➡️<a href='%s'>перейти</a>", link)

	template = template + "\nСтатус:" + lot.LotStatus + linkStr

	err := rest.SendTgMessage(context.Background(), config.Instance.Tg.Url, config.Instance.Tg.Token, template)

	if err != nil {
		logging.Log().Errorln(err)
	}
}
