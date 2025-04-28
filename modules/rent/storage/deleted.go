package storage

import "context"

func (s *SqlModel) DeletedRent(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("rent").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
