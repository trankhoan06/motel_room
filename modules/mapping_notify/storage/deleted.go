package storage

import (
	"context"
)

func (s *SqlModel) DeletedMappingNotify(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("mapping_notify").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
