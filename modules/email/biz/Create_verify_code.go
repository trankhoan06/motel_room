package biz

import (
	"context"
	"errors"
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
	_, err1 := biz.store.FindCodeVerify(ctx, map[string]interface{}{"email": email})
	if err1 == nil {
		return nil, errors.New("row has available")
	}

	var verifyEmail model.CreateVerifyAccount
	verifyEmail.Email = email
	verifyEmail.Verify = false
	typeCode := model.TypeVerifyEmail
	verifyEmail.Type = &typeCode
	verifyEmail.Code = common.GenerateRandomCode()
	now := time.Now().Add(-7 * time.Hour)
	verifyEmail.Expire = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.CreateCodeVerify(ctx, &verifyEmail); err != nil {
		return nil, err
	}
	return &verifyEmail, nil
}
