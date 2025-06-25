package biz

import (
	"context"
	"fmt"
	"main.go/modules/user_like_room/model"
)

func (biz *UserLikeRoomCommon) NewListUserLikeRoom(ctx context.Context, rentId int) (*[]model.UserLikeRoom, error) {
	fmt.Println(rentId)
	data, err := biz.userLikeRoom.ListUserLikeRoom(ctx, map[string]interface{}{"rent_id": rentId})
	if err != nil {
		return nil, err
	}
	return data, nil
}
