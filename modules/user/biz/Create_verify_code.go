package biz

import (
	"context"
	"main.go/common"
	emailSend "main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (biz *UserCommonBiz) NewCreateVerifyCodeEmail(ctx context.Context, verify *model.VerifyAccountCode, expire int) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": verify.Email})
	if err != nil {
		return common.ErrEmailNoExist(err)
	}
	var verifyEmail model.CreateVerifyAccount
	verifyEmail.UserId = user.Id
	verifyEmail.IsVerifyEmail = verify.IsVerifyEmail
	verifyEmail.IsForgotPassword = verify.IsForgotPassword
	verifyEmail.Code = common.GenerateRandomCode()
	now := time.Now().Add(-7 * time.Hour)
	verifyEmail.Expire = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.CreateCodeVerify(ctx, &verifyEmail); err != nil {
		return err
	}
	emailSend.SendVerifyEmail(user.Email, verifyEmail.Code)
	return nil
}
