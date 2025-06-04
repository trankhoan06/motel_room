package storage

import (
	"context"
	"main.go/modules/search/model"
)

func (s *SqlModel) CreateSearch(ctx context.Context, data *model.CreateSearch) error {
	db := s.db.Begin()
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
