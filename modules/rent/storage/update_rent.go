package storage

import (
	"context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) UpdateRent(ctx context.Context, data *model.UpdateRent) error {
	if err := s.db.Where("id=?", data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
