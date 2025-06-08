package storage

import (
	"context"
)

func (s *SqlModel) DeletedUserLikeRoom(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("user_like_room").Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
