package storage

import (
	"context"
	"main.go/modules/user_like_room/model"
)

func (s *SqlModel) ListUserLikeRoom(ctx context.Context, cond map[string]interface{}) (*[]model.UserLikeRoom, error) {
	var data []model.UserLikeRoom
	db := s.db.Preload("SimpleRent")
	if err := db.Table("user_like_room").Where(cond).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
