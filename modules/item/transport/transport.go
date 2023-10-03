package transport

import "gorm.io/gorm"

type Transport struct {
	db *gorm.DB
}

func NewTransport(db *gorm.DB) *Transport {
	return &Transport{db: db}
}
