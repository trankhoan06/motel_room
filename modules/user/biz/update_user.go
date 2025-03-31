package biz

import (
	"context"
	"main.go/common"
	"main.go/modules/user/model"
)

func (biz *UserCommonBiz) NewUpdateUser(ctx context.Context, data *model.UpdateUser) error {
	if data.Email == "" {
		return common.ErrLogin
	}
	if err := biz.store.UpdateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
