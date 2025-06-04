package storage

import (
	"context"
	"main.go/modules/room_reviews/model"
)

func (s *SqlModel) CreateReview(ctx context.Context, data *model.CreateReviews) error {
	db := s.db.Begin()
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		return err
	}
	return nil
}
