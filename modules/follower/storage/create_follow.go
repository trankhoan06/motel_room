package storage

import (
	"context"
	"main.go/modules/follower/model"
)

func (s *SqlModel) CreateFollow(ctx context.Context, data *model.CreateFollower) error {
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
