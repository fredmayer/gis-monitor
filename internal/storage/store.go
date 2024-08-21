package storage

import (
	"context"
	"errors"
	"gis-crawler/internal/config"
	"gis-crawler/internal/models"
	"gis-crawler/internal/storage/mysql"
	"gis-crawler/pkg/logging"
	gsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func MustLoad(ctx context.Context, cfg *config.Config) *Store {

	mdb := mysql.NewClient(ctx, mysql.Options{
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password, cfg.Db.Name,
	})
	gormDB, err := gorm.Open(gsql.New(gsql.Config{
		Conn: mdb,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = gormDB.AutoMigrate(&models.Lot{})
	if err != nil {
		logging.Get().Error(err)
	}

	return &Store{
		db: gormDB,
	}

}

func (s *Store) GetByUID(uid string) (*models.Lot, bool) {
	//TODO implement me
	var lot models.Lot
	err := s.db.First(&lot, "id = ?", uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &lot, false
	}
	if err != nil {
		logging.Get().Error(err)
		return &lot, false
	}

	return &lot, true
}

func (s *Store) Save(lot *models.Lot) error {
	result := s.db.Create(&lot)
	return result.Error
}

func (s *Store) Update(lot *models.Lot) error {
	result := s.db.Save(&lot)
	return result.Error
}
