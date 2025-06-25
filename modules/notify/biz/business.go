package biz

import (
	"context"
	"main.go/modules/notify/model"
	"main.go/worker"
)

type NotifyBiz interface {
	CreateNotify(ctx context.Context, data *model.CreateNotify) error
	DeletedNotify(ctx context.Context, cond map[string]interface{}) error
	FindNotify(ctx context.Context, cond map[string]interface{}) (*model.Notify, error)
	UpdateNotify(ctx context.Context, data *model.UpdateNotify) error
}
type NotifyMappingBiz interface {
	DeletedMappingNotify(ctx context.Context, cond map[string]interface{}) error
}
type NotifyTaskCommon struct {
	Notify          NotifyBiz
	taskDistributor worker.TaskDistributor
}

func NewNotifyTaskCommon(notify NotifyBiz, distributor worker.TaskDistributor) *NotifyTaskCommon {
	return &NotifyTaskCommon{notify, distributor}
}

type NotifyCommon struct {
	Notify  NotifyBiz
	mapping NotifyMappingBiz
}

func NewNotifyCommon(notify NotifyBiz, mapping NotifyMappingBiz) *NotifyCommon {
	return &NotifyCommon{notify, mapping}
}
