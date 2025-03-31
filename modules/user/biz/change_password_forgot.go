package biz

import (
	"context"
	"main.go/common"
	"main.go/modules/user/model"
	"time"
)

func (biz *RegisterUserBiz) NewChangePasswordForgot(ctx context.Context, password *model.NewPasswordForgot) error {
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
	if !v.Verify {
		return common.ErrVerify
	}
	now := time.Now().Add(-7 * time.Hour)
	if v.Expire.Before(now) {
		return common.ErrForgotPassword
	}
	password.NewPassword = biz.hash.Hash(user.Salt + password.NewPassword)
	if err := biz.store.ChangePassword(ctx, user.Id, password.NewPassword); err != nil {
		return err
	}
	if err := biz.store.UpdateVerifyCode(ctx, map[string]interface{}{"user_id": user.Id}, map[string]interface{}{"expire": now}); err != nil {
		return err
	}

	return nil
}
