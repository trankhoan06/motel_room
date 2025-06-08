package biz

import (
	"context"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/user_like_room/model"
)

type UserLikeRoomBiz interface {
	CreateUserLikeRoom(ctx context.Context, data *model.CreateUserLikeRoom) error
	DeletedUserLikeRoom(ctx context.Context, cond map[string]interface{}) error
	ListUserLikeRoom(ctx context.Context, cond map[string]interface{}) (*[]model.UserLikeRoom, error)
	FindUserLikeRoom(ctx context.Context, cond map[string]interface{}) (*model.UserLikeRoom, error)
}
type RentBiz interface {
	FindRent(ctx context.Context, cond map[string]interface{}) (*modelRent.Rent, error)
}
type UserLikeRoomCommon struct {
	userLikeRoom UserLikeRoomBiz
	rent         RentBiz
}

func NewUserLikeRoomCommon(userLikeRoom UserLikeRoomBiz, rent RentBiz) *UserLikeRoomCommon {
	return &UserLikeRoomCommon{userLikeRoom: userLikeRoom, rent: rent}
}
