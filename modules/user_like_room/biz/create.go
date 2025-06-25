package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/user_like_room/model"
)

func (biz *UserLikeRoomCommon) NewCreateUserLikeRoom(ctx context.Context, data *model.CreateUserLikeRoom) error {
	if data.RentId == 0 {
		return common.ErrRequire(errors.New("rent id is require"))
	}
	if _, err := biz.userLikeRoom.FindUserLikeRoom(ctx, map[string]interface{}{"user_id": data.UserId, "rent_id": data.RentId}); err == nil {
		return errors.New("user has been liked rent")
	}
	_, err := biz.rent.FindRent(ctx, map[string]interface{}{"id": data.RentId})
	if err != nil {
		return err
	}
	if err := biz.userLikeRoom.CreateUserLikeRoom(ctx, data); err != nil {
		return err
	}
	return nil
}
