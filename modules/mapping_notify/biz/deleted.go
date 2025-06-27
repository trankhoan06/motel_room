package biz

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

func (biz *MappingNotifyCommon) NewDeletedMappingNotify(ctx context.Context, data *model.CreateMappingNotify) error {
	if _, err := biz.mapping.FindMappingNotify(ctx, map[string]interface{}{"notification_id": data.NotificationId, "user_id": data.UserId}); err != nil {
		return err
	}
	if err := biz.mapping.DeletedMappingNotify(ctx, map[string]interface{}{"notification_id": data.NotificationId, "user_id": data.UserId}); err != nil {
		return err
	}
	return nil

}
