package storage

import (
	"context"
	"main.go/modules/room_reviews/model"
)

func (s *SqlModel) UpdateReview(ctx context.Context, data *model.UpdateReviews) error {
	if err := s.db.Updates(data).Error; err != nil {
		return err
	}
	return nil
}
