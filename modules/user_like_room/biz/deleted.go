package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/user_like_room/model"
)

func (biz *UserLikeRoomCommon) NewDeletedUserLikeRoom(ctx context.Context, data *model.CreateUserLikeRoom) error {
	if data.RentId == 0 {
		return common.ErrRequire(errors.New("rent id is require"))
	}
	if _, err := biz.userLikeRoom.FindUserLikeRoom(ctx, map[string]interface{}{"user_id": data.UserId, "rent_id": data.RentId}); err != nil {
		return err
	}
	_, err := biz.rent.FindRent(ctx, map[string]interface{}{"id": data.RentId})
	if err != nil {
		return err
	}
	if err := biz.userLikeRoom.DeletedUserLikeRoom(ctx, map[string]interface{}{"user_id": data.UserId, "rent_id": data.RentId}); err != nil {
		return err
	}
	return nil
}
