package storage

import (
	"context"
	"main.go/modules/notify/model"
)

func (s *SqlModel) UpdateNotify(ctx context.Context, data *model.UpdateNotify) error {
	if err := s.db.Where("id=?", data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
