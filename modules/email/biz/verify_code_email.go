package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/modules/email/model"
	"time"
)

func (biz LoginBiz) NewVerifyEmail(ctx context.Context, verify *model.VerifyAccountCode, expire int) (*tokenprovider.Token, error) {
	if verify.Email == "" {
		return nil, common.ErrEmailRequire
	}
	user, err := biz.user.FindUser(ctx, map[string]interface{}{"email": verify.Email})
	if err != nil {
		return nil, common.ErrEmailNoExist(err)
	}
	v, errVerify := biz.store.FindCodeVerify(ctx, map[string]interface{}{"email": user.Email, "type": int(model.TypeVerifyEmail)})
	if v.Verify {
		return nil, errors.New("code has been verified")
	}
	if errVerify != nil {
		return nil, err
	}
	now := time.Now().Add(-7 * time.Hour)
	if v.Expire.Before(now) {
		return nil, common.ErrVerifyCodeExpire
	}
	if v.Code != verify.Code {
		return nil, common.ErrVerifyCode
	}
	if err := biz.store.UpdateVerifyCode(ctx, map[string]interface{}{"id": user.Id}, map[string]interface{}{"verify": 1}); err != nil {
		return nil, err
	}
	if err := biz.store.UpdateVerifyEmail(ctx, map[string]interface{}{"id": user.Id}); err != nil {
		return nil, err
	}
	var payload = &common.Payload{
		UId:  user.Id,
		Role: user.Role,
	}
	token, errToken := biz.provider.Generate(payload, expire)
	if errToken != nil {
		return nil, errToken
	}
	return &token, nil
}
