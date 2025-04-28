package biz

import (
	"context"
	"errors"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (biz *SendEMailBiz) NewForgotPassword(ctx context.Context, data *model.ForgotPassword, expire int) (*model.VerifyToken, error) {
	if data.Email == "" {
		return nil, common.ErrEmailRequire
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, common.ErrEmailNoExist(errors.New("user don't exist"))
	}
	var verifyEmail model.CreateVerifyAccount
	verifyEmail.UserId = user.Id
	verifyEmail.Token = common.GetSalt(30)
	verifyEmail.Code = common.GenerateRandomCode()
	now := time.Now().Add(-7 * time.Hour)
	verifyEmail.Expire = now.Add(time.Duration(expire) * time.Second)
	if err1 := biz.store.CreateCodeVerify(ctx, &verifyEmail); err1 != nil {
		return nil, err1
	}
	emailSend.SendForgotPassword(user.Email, verifyEmail.Code, biz.cfg)
	return &model.VerifyToken{
		Token:   verifyEmail.Token,
		Email:   data.Email,
		IsLogin: false,
	}, nil
}
