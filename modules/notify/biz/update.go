package biz

import (
	"context"
	"errors"
	"main.go/modules/notify/model"
)

func (biz *NotifyCommon) NewUpdateNotify(ctx context.Context, data *model.UpdateNotify) error {
	notify, err := biz.Notify.FindNotify(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}
	if notify.UserId != data.Id {
		return errors.New("you can't deleted a notification that you did not create")
	}
	if err := biz.Notify.UpdateNotify(ctx, data); err != nil {
		return err
	}
	return nil
}
