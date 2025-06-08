package storage

import (
	"context"
	"main.go/modules/user_like_room/model"
)

func (s *SqlModel) CreateUserLikeRoom(ctx context.Context, data *model.CreateUserLikeRoom) error {
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
