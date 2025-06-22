package biz

import (
	"context"
	"main.go/modules/user/model"
)

func (biz *UserCommonBiz) NewChangeStatusAccount(ctx context.Context, cond map[string]interface{}, status *model.StatusUser) error {
	if err := biz.store.ChangeStatusUser(ctx, cond, status); err != nil {
		return err
	}
	return nil
}
