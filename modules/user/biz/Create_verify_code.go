package biz

import (
	"context"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (biz *SendEMailBiz) NewCreateVerifyCodeEmail(ctx context.Context, verify *model.VerifyAccountCode, expire int) (*model.VerifyToken, error) {
	if verify.Email == "" {
		return nil, common.ErrEmailRequire
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": verify.Email})
	if err != nil {
		return nil, common.ErrEmailNoExist(err)
	}
	var verifyEmail model.CreateVerifyAccount
	verifyEmail.UserId = user.Id
	verifyEmail.Token = common.GetSalt(30)
	verifyEmail.Code = common.GenerateRandomCode()
	now := time.Now().Add(-7 * time.Hour)
	verifyEmail.Expire = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.CreateCodeVerify(ctx, &verifyEmail); err != nil {
		return nil, err
	}
	emailSend.SendVerifyEmail(user.Email, verifyEmail.Code, biz.cfg)
	return &model.VerifyToken{
		Token:   verifyEmail.Token,
		Email:   verify.Email,
		IsLogin: false,
	}, nil
}
