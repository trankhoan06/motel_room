package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) ChangeStatusUser(ctx context.Context, cond map[string]interface{}, status *model.StatusUser) error {
	if err := s.db.Where(cond).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
