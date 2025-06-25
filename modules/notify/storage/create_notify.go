package storage

import (
	"context"
	"main.go/modules/notify/model"
)

func (s *SqlModel) CreateNotify(ctx context.Context, data *model.CreateNotify) error {
	db := s.db.Begin()
	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit(); err != nil {
		db.Rollback()
		return err.Error
	}
	return nil
}
