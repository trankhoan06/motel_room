package storage

import "context"

func (s *SqlModel) DeletedReview(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("room_reviews").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
