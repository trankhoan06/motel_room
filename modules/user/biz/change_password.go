package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/user/model"
)

func (biz *RegisterUserBiz) NewChangePassword(ctx context.Context, data *model.ChangePassword) error {
	if data.Email == "" {
		return common.ErrLogin
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}
	data.OldPassword = biz.hash.Hash(user.Salt + data.OldPassword)
	if data.OldPassword != user.Password {
		return errors.New("Old password error")
	}
	data.NewPassword = biz.hash.Hash(user.Salt + data.NewPassword)
	if errChange := biz.store.ChangePassword(ctx, data.Id, data.NewPassword); errChange != nil {
		return err
	}
	return nil
}
