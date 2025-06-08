package storage

import (
	"context"
	"main.go/modules/user_like_room/model"
)

func (s *SqlModel) FindUserLikeRoom(ctx context.Context, cond map[string]interface{}) (*model.UserLikeRoom, error) {
	var data model.UserLikeRoom
	db := s.db.Where("status=?", model.StatusDoingUserLikeRoom).Preload("SimpleRent")
	if err := db.Table("user_like_room").Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
