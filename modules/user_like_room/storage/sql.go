package storage

import "gorm.io/gorm"

type SqlModel struct {
	db *gorm.DB
}

func NewSqlMode(db *gorm.DB) *SqlModel {
	return &SqlModel{db: db}
}
