package storage

import (
	"context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) ListRent(ctx context.Context, cond map[string]interface{}) (*[]model.Rent, error) {
	var data []model.Rent
	db := s.db.Preload("Owner").Where("status <> ?", model.StatusDeletedRent)
	if err := db.Table("rent").Where(cond).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
