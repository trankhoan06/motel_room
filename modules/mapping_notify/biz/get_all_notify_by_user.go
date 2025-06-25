package biz

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

func (biz *MappingNotifyCommon) NewGetAllNotifyByUser(ctx context.Context, userId int) (*[]model.MappingNotify, error) {
	res, err := biz.mapping.GetAllNotifyByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
