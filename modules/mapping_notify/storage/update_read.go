package storage

import (
	"context"
)

func (s *SqlModel) UpdateRead(ctx context.Context, cond map[string]interface{}, read bool) error {
	if err := s.db.Table("mapping_notify").Where(cond).Update("is_read", read).Error; err != nil {
		return err
	}
	return nil
}
