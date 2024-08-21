package gis

import "gis-crawler/internal/models"

type StoreInterface interface {
	GetByUID(uid string) (*models.Lot, bool)
	Save(lot *models.Lot) error
	Update(lot *models.Lot) error
}

type ClientInterface interface {
	CreateRequest(method string, path string)
	AddParam(key string, value string)
	Send() []byte
}
