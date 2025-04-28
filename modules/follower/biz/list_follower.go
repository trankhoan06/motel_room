package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/follower/model"
	modelUser "main.go/modules/user/model"
)

func (biz *FollowUserCommon) NewListFollower(ctx context.Context, userId int) (*[]model.Follower, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, common.ErrUserExist(err)
	}
	if *user.Status == modelUser.StatusUserDeleted {
		return nil, common.ErrUserExist(errors.New("err user"))
	}
	follows, errFolow := biz.store.ListFollow(ctx, map[string]interface{}{"user_id": userId})
	if errFolow != nil {
		return nil, errFolow
	}
	return follows, nil
}
