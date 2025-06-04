package storage

import (
	"context"
	"main.go/modules/room_reviews/model"
)

func (s *SqlModel) FindReview(ctx context.Context, cond map[string]interface{}) (*model.Reviews, error) {
	var data model.Reviews
	db := s.db.Preload("User")
	if err := db.Table("room_reviews").Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
