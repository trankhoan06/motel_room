package storage

import (
	"context"
)

func (s *SqlModel) DeletedAfter7Day(ctx context.Context) error {
	if err := s.db.Table("rent").Where("update_at < NOW() - INTERVAL 6 DAY").Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
