package biz

import (
	"context"
	"main.go/common"
	"main.go/modules/user/model"
	"time"
)

func (biz *UserCommonBiz) NewVerifyForgotPassword(ctx context.Context, password *model.VerifyAccountCode, expire int) error {
	if password.Email == "" {
		return common.ErrEmailRequire
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": password.Email})
	if err != nil {
		return err
	}
	v, errVerify := biz.store.FindCodeVerify(ctx, map[string]interface{}{"user_id": user.Id, "token": password.Token})
	if errVerify != nil {
		return errVerify
	}
	if v.Code != password.Code {
		return common.ErrVerifyCode
	}
	now := time.Now()
	now = now.Add(-7 * time.Hour)
	if v.Expire.Before(now) {
		return common.ErrVerifyCodeExpire
	}
	now = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.UpdateVerifyCode(ctx, map[string]interface{}{"user_id": user.Id}, map[string]interface{}{"verify": 1, "expire": now}); err != nil {
		return err
	}
	return nil
}
