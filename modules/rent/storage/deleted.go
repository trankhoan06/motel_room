package storage

import (
	"context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) DeletedRent(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("rent").Where(cond).Update("status", model.StatusDeletedRent).Error; err != nil {
		return err
	}
	return nil
}
