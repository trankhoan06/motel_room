package biz

import (
	"context"
	"errors"
	"fmt"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (biz *RegisterUserBiz) NewRegister(ctx context.Context, data *model.Register, expiry int) (*model.VerifyToken, error) {

	if data.Email == "" {
		return nil, errors.New("email is request")
	}
	if _, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email}); err == nil {
		return nil, errors.New("user already exists")
	}
	data.Salt = common.GetSalt(50)
	data.Password = biz.hash.Hash(data.Salt + data.Password)
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return nil, errors.New("error create user")
	}
	var verifyEmail model.CreateVerifyAccount
	verifyEmail.UserId = data.Id
	verifyEmail.Code = common.GenerateRandomCode()
	now := time.Now().Add(-7 * time.Hour)
	verifyEmail.Expire = now.Add(time.Duration(expiry) * time.Second)
	if err := biz.store.CreateCodeVerify(ctx, &verifyEmail); err != nil {
		fmt.Print(err)
	}
	emailSend.SendVerifyEmail(data.Email, verifyEmail.Code, biz.cfg)
	return &model.VerifyToken{
		Token:   verifyEmail.Token,
		Email:   data.Email,
		IsLogin: true,
	}, nil
}
