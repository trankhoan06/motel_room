package storage

import (
	"context"
)

func (s *SqlModel) DeletedSearch(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("search").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
