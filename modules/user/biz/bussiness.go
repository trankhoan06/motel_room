package biz

import (
	"context"
	"main.go/component/tokenprovider"
	"main.go/config"
	"main.go/modules/user/model"
)

type UserBiz interface {
	UpdateUser(ctx context.Context, data *model.UpdateUser) error
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	CreateUser(ctx context.Context, data *model.Register) error
	DeletedUser(ctx context.Context, cond map[string]interface{}) error
	CreateCodeVerify(ctx context.Context, data *model.CreateVerifyAccount) error
	FindCodeVerify(ctx context.Context, cond map[string]interface{}) (*model.VerifyAccount, error)
	UpdateVerifyCode(ctx context.Context, cond map[string]interface{}, update map[string]interface{}) error
	UpdateVerifyEmail(ctx context.Context, cond map[string]interface{}) error
	ChangePassword(ctx context.Context, id int, password string) error
}
type UserCommonBiz struct {
	store UserBiz
}

func NewUserCommonBiz(store UserBiz) *UserCommonBiz {
	return &UserCommonBiz{store: store}
}

type SendEMailBiz struct {
	store UserBiz
	cfg   *config.Config
}

func NewSendEmailBiz(store UserBiz, cfg *config.Config) *SendEMailBiz {
	return &SendEMailBiz{store: store, cfg: cfg}
}

type Hasher interface {
	Hash(str string) string
}
type RegisterUserBiz struct {
	store UserBiz
	hash  Hasher
	cfg   *config.Config
}

func NewRegisterUserBiz(store UserBiz, hash Hasher, cfg *config.Config) *RegisterUserBiz {
	return &RegisterUserBiz{store: store, hash: hash, cfg: cfg}
}

type LoginBiz struct {
	store    UserBiz
	provider tokenprovider.TokenProvider
	hash     Hasher
	cfg      *config.Config
}

func NewLoginBiz(store UserBiz, provider tokenprovider.TokenProvider, hash Hasher, cfg *config.Config) *LoginBiz {
	return &LoginBiz{store: store, provider: provider, hash: hash, cfg: cfg}
}
