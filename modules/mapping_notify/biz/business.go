package biz

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

type MappingNotifyBiz interface {
	CreateMappingNotify(ctx context.Context, data *[]model.CreateMappingNotify, size int) error
	DeletedMappingNotify(ctx context.Context, cond map[string]interface{}) error
	FindMappingNotify(ctx context.Context, cond map[string]interface{}) (*model.MappingNotify, error)
	GetAllNotifyByUser(ctx context.Context, userId int) (*[]model.MappingNotify, error)
	UpdateRead(ctx context.Context, cond map[string]interface{}, read bool) error
}
type MappingNotifyCommon struct {
	mapping MappingNotifyBiz
}

func NewMappingNotifyCommon(mapping MappingNotifyBiz) *MappingNotifyCommon {
	return &MappingNotifyCommon{mapping}
}
