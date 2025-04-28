package storage

import (
	"context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) CreateRent(ctx context.Context, data *model.CreateRent) error {
	db := s.db.Begin()
	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
