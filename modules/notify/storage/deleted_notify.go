package storage

import (
	"context"
)

func (s *SqlModel) DeletedNotify(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("notify").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
