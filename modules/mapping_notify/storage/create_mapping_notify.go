package storage

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

func (s *SqlModel) CreateMappingNotify(ctx context.Context, data *[]model.CreateMappingNotify, size int) error {
	db := s.db.Begin()
	if err := db.CreateInBatches(&data, size).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit(); err != nil {
		db.Rollback()
		return err.Error
	}
	return nil
}
