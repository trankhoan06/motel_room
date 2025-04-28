package biz

import (
	"context"
	"main.go/modules/follower/model"
	modelUser "main.go/modules/user/model"
)

type FollowerBiz interface {
	CreateFollow(ctx context.Context, data *model.CreateFollower) error
	CancelFollow(ctx context.Context, cond map[string]interface{}) error
	ListFollow(ctx context.Context, cond map[string]interface{}) (*[]model.Follower, error)
	FindFollower(ctx context.Context, cond map[string]interface{}) (*model.Follower, error)
}
type UserBiz interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
	UpdateFollow(ctx context.Context, cond map[string]interface{}, Expr, calculation string) error
}
type FollowCommon struct {
	store FollowerBiz
}

func NewFollowCommon(store FollowerBiz) *FollowCommon {
	return &FollowCommon{
		store: store,
	}
}

type FollowUserCommon struct {
	store     FollowerBiz
	storeUser UserBiz
}

func NewFollowUserCommon(store FollowerBiz, StoreUser UserBiz) *FollowUserCommon {
	return &FollowUserCommon{
		store:     store,
		storeUser: StoreUser,
	}
}
