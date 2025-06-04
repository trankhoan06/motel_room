package storage

import (
	"context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) UpdateRate(ctx context.Context, data *model.RateRent) error {
	if err := s.db.Table("rent").Where("id=?", data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
