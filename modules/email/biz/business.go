package biz

import (
	"context"
	"main.go/component/tokenprovider"
	"main.go/modules/email/model"
	modelUser "main.go/modules/user/model"
)

type EmailBiz interface {
	CreateCodeVerify(ctx context.Context, data *model.CreateVerifyAccount) error
	FindCodeVerify(ctx context.Context, cond map[string]interface{}) (*model.VerifyAccount, error)
	UpdateVerifyCode(ctx context.Context, cond map[string]interface{}, update map[string]interface{}) error
	UpdateVerifyEmail(ctx context.Context, cond map[string]interface{}) error
	UpdateSendCodeEmail(ctx context.Context, data *model.CreateVerifyAccount) error
}
type UserEmailBiz interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}
type UserCommonBiz struct {
	store EmailBiz
}

func NewUserCommonBiz(store EmailBiz) *UserCommonBiz {
	return &UserCommonBiz{store: store}
}

type SendEMailBiz struct {
	store EmailBiz
	user  UserEmailBiz
}

func NewSendEmailBiz(store EmailBiz, user UserEmailBiz) *SendEMailBiz {
	return &SendEMailBiz{store: store, user: user}
}

type Hasher interface {
	Hash(str string) string
}
type RegisterEmailBiz struct {
	store EmailBiz
	hash  Hasher
}

func NewRegisterEmailBiz(store EmailBiz, hash Hasher) *RegisterEmailBiz {
	return &RegisterEmailBiz{store: store, hash: hash}
}

type LoginBiz struct {
	store    EmailBiz
	user     UserEmailBiz
	provider tokenprovider.TokenProvider
	hash     Hasher
}

func NewLoginBiz(store EmailBiz, provider tokenprovider.TokenProvider, hash Hasher) *LoginBiz {
	return &LoginBiz{store: store, provider: provider, hash: hash}
}
