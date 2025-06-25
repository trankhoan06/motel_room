package biz

import (
	"context"
	"errors"
	"main.go/modules/notify/model"
)

func (biz *NotifyCommon) NewDeletedNotify(ctx context.Context, data *model.DeletedNotify) error {
	notify, err := biz.Notify.FindNotify(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}
	if notify.UserId != data.Id {
		return errors.New("you can't deleted a notification that you did not create")
	}
	if err := biz.Notify.DeletedNotify(ctx, map[string]interface{}{"id": data.Id}); err != nil {
		return err
	}
	if err := biz.mapping.DeletedMappingNotify(ctx, map[string]interface{}{"notification_id": data.Id}); err != nil {
		return err
	}

	return nil
}
