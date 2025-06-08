package biz

import (
	"context"
	"main.go/common"
	"main.go/modules/email/model"
	"time"
)

func (biz *SendEMailBiz) NewCreateVerifyCodeEmail(ctx context.Context, email string, expire int) (*model.CreateVerifyAccount, error) {
	//check user exist
	_, err := biz.user.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, common.ErrEmailNoExist(err)
	}

	var verifyEmail model.CreateVerifyAccount
	verifyEmail.Email = email
	verifyEmail.Verify = false
	*verifyEmail.Type = model.TypeVerifyEmail
	verifyEmail.Code = common.GenerateRandomCode()
	now := time.Now().Add(-7 * time.Hour)
	verifyEmail.Expire = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.CreateCodeVerify(ctx, &verifyEmail); err != nil {
		return nil, err
	}
	return &verifyEmail, nil
}
