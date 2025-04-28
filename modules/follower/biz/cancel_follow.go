package biz

import (
	"context"
	"main.go/common"
	"main.go/modules/follower/model"
)

func (biz *FollowUserCommon) NewCancelFollow(ctx context.Context, data *model.CreateFollower) error {
	if _, err := biz.store.FindFollower(ctx, map[string]interface{}{"user_id": data.UserId, "follower_id": data.FollowerId}); err != nil {
		return common.ErrCancel(err)
	}
	if err := biz.storeUser.UpdateFollow(ctx, map[string]interface{}{"id": data.FollowerId}, "amount_following", "-"); err != nil {
		return err
	}
	if err := biz.storeUser.UpdateFollow(ctx, map[string]interface{}{"id": data.UserId}, "amount_follower", "-"); err != nil {
		return err
	}
	if err := biz.store.CancelFollow(ctx, map[string]interface{}{"user_id": data.UserId, "follower_id": data.FollowerId}); err != nil {
		return err
	}
	return nil
}
