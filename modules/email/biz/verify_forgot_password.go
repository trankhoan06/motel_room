package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/email/model"
	"time"
)

func (biz SendEMailBiz) NewVerifyForgotPassword(ctx context.Context, password *model.VerifyAccountCode, expire int) error {
	if password.Email == "" {
		return common.ErrEmailRequire
	}
	user, err := biz.user.FindUser(ctx, map[string]interface{}{"email": password.Email})
	if err != nil {
		return err
	}
	v, errVerify := biz.store.FindCodeVerify(ctx, map[string]interface{}{"email": user.Email, "type": int(model.TypeForgotPassword)})
	if errVerify != nil {
		return errVerify
	}
	if v.Verify {
		return errors.New("code has been verified")
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
	if err := biz.store.UpdateVerifyCode(ctx, map[string]interface{}{"email": user.Email}, map[string]interface{}{"verify": 1, "expire": now}); err != nil {
		return err
	}
	return nil
}
