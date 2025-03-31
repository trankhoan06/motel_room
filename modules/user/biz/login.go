package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/modules/user/model"
)

func (biz *LoginBiz) NewLogin(ctx context.Context, data *model.Login, expiry int) (tokenprovider.Token, error) {
	if data.Email == "" {
		return nil, common.ErrEmailRequire
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, errors.New("email don't exist")
	}
	if *user.Status == model.StatusUserDeleted {
		return nil, errors.New("user is deleted")
	}
	pass := biz.hash.Hash(user.Salt + data.Password)
	if pass != user.Password {
		return nil, errors.New("account or password wrong")
	}
	if !user.IsEMail {
		return nil, common.ErrVerifyEmail
	}
	var payload = &common.Payload{
		UId:  user.Id,
		Role: user.Role,
	}
	token, errToken := biz.provider.Generate(payload, expiry)
	if errToken != nil {
		return nil, errToken
	}
	return token, nil
}
