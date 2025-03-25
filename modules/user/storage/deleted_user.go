package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) DeletedUser(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Where(cond).Update("status", model.StatusUserDeleted).Error; err != nil {
		return err
	}
	return nil
}
