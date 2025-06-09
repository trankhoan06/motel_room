package biz

import (
	"context"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/config"
	modelEmail "main.go/modules/email/model"
	"main.go/modules/user/model"
	"main.go/worker"
)

type UserBiz interface {
	UpdateUser(ctx context.Context, data *model.UpdateUser) error
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	CreateUser(ctx context.Context, data *model.Register) error
	DeletedUser(ctx context.Context, cond map[string]interface{}) error
	//UpdateVerifyEmail(ctx context.Context, cond map[string]interface{}) error
	ChangePassword(ctx context.Context, id int, password string) error
}
type UserEmailBiz interface {
	FindCodeVerify(ctx context.Context, cond map[string]interface{}) (*modelEmail.VerifyAccount, error)
	UpdateVerifyCode(ctx context.Context, cond map[string]interface{}, update map[string]interface{}) error
	CreateCodeVerify(ctx context.Context, data *modelEmail.CreateVerifyAccount) error
}
type UserCommonBiz struct {
	store UserBiz
}

func NewUserCommonBiz(store UserBiz) *UserCommonBiz {
	return &UserCommonBiz{store: store}
}

type DeletedRentCase interface {
	NewDeletedRent(ctx context.Context, com *common.IdCommon) error
}
type DeletedUserCommonBiz struct {
	store UserBiz
	rent  DeletedRentCase
}

func NewDeletedUserCommonBiz(store UserBiz, rent DeletedRentCase) *DeletedUserCommonBiz {
	return &DeletedUserCommonBiz{store: store, rent: rent}
}

type Hasher interface {
	Hash(str string) string
}
type RegisterUserBiz struct {
	store           UserBiz
	hash            Hasher
	email           UserEmailBiz
	taskDistributor worker.TaskDistributor
}

func NewRegisterUserBiz(store UserBiz, hash Hasher, email UserEmailBiz, taskDistributor worker.TaskDistributor) *RegisterUserBiz {
	return &RegisterUserBiz{store: store, hash: hash, email: email, taskDistributor: taskDistributor}
}

type RegisterEmailUserBiz struct {
	store UserBiz
	email UserEmailBiz
}

func NewRegisterEmailUserBiz(store UserBiz, email UserEmailBiz) *RegisterEmailUserBiz {
	return &RegisterEmailUserBiz{store: store, email: email}
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
