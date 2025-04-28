package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/follower/model"
	modelUser "main.go/modules/user/model"
)

func (biz *FollowUserCommon) NewCreateFollow(ctx context.Context, data *model.CreateFollower) error {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"id": data.UserId})
	if err != nil {
		return common.ErrUserExist(errors.New("account has been deleted or not exist"))
	}
	if *user.Status == modelUser.StatusUserDeleted {
		return common.ErrUserExist(errors.New("account has been deleted or not exist"))
	}
	_, err1 := biz.store.FindFollower(ctx, map[string]interface{}{"user_id": data.UserId, "follower_id": data.FollowerId})
	if err1 == nil {
		return common.ErrFollow(errors.New("err follow"))
	}
	if err := biz.storeUser.UpdateFollow(ctx, map[string]interface{}{"id": data.FollowerId}, "amount_following", "+"); err != nil {
		return err
	}
	if err := biz.storeUser.UpdateFollow(ctx, map[string]interface{}{"id": data.UserId}, "amount_follower", "+"); err != nil {
		return err
	}
	errCreate := biz.store.CreateFollow(ctx, data)
	if errCreate != nil {
		return errCreate
	}
	return nil
}
