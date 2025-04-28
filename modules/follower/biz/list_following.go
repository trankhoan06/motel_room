package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/follower/model"
	modelUser "main.go/modules/user/model"
)

func (biz *FollowUserCommon) NewListFollowing(ctx context.Context, userId int) (*[]model.Follower, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, common.ErrUserExist(err)
	}
	if *user.Status == modelUser.StatusUserDeleted {
		return nil, common.ErrUserExist(errors.New("err user"))
	}
	follows, errFollow := biz.store.ListFollow(ctx, map[string]interface{}{"follower_id": userId})
	if errFollow != nil {
		return nil, errFollow
	}
	return follows, nil
}
