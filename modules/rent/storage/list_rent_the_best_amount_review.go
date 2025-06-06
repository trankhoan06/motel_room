package storage

import (
	"context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) ListRentTheBestAmountReview(ctx context.Context, limit int) (*[]model.Rent, error) {
	db := s.db.Preload("Owner").Where("status <> ?", model.StatusDeletedRent)
	var data []model.Rent
	if err := db.Table("rent").Order("amount_rate DESC").Limit(limit).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil

}
