package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) CreateUser(ctx context.Context, data *model.Register) error {
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
