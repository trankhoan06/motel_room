package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) UpdateUser(ctx context.Context, data *model.UpdateUser) error {
	db := s.db.Begin()
	if err := db.Where("email=?", data.Email).Updates(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
