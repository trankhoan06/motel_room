package biz

import (
	"context"
	"main.go/modules/user_like_room/model"
)

func (biz *UserLikeRoomCommon) NewListUserLikeRoom(ctx context.Context, userId int) (*[]model.UserLikeRoom, error) {
	data, err := biz.userLikeRoom.ListUserLikeRoom(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}
	return data, nil
}
