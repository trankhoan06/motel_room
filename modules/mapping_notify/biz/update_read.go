package biz

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

func (biz *MappingNotifyCommon) NewUpdateRead(ctx context.Context, data *model.CreateMappingNotify) error {
	if _, err := biz.mapping.FindMappingNotify(ctx, map[string]interface{}{"notification_id": data.NotificationId, "user_id": data.UserId}); err != nil {
		return err
	}
	if err := biz.mapping.UpdateRead(ctx, map[string]interface{}{"notification_id": data.NotificationId, "user_id": data.UserId}, true); err != nil {
		return err
	}
	return nil
}
