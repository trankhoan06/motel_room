package storage

import (
	"context"
)

func (s *SqlModel) CancelFollow(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("follow").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
