package biz

import (
	"context"
	"github.com/hibiken/asynq"
	"main.go/modules/notify/model"
	"main.go/worker"
	"time"
)

func (biz *NotifyTaskCommon) NewCreateNotify(ctx context.Context, data *model.CreateNotify) error {
	if err := biz.Notify.CreateNotify(ctx, data); err != nil {
		return err
	}
	//task

	taskPayload := worker.PayloadCreateNotify{
		UserId:         data.UserId,
		NotificationId: data.Id,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendResetCodePassword),
	}
	_ = biz.taskDistributor.DistributeTaskCreateNotify(ctx, &taskPayload, opts...)
	return nil
}
