package biz

import (
	"context"
	"errors"
	"github.com/hibiken/asynq"
	"main.go/common"
	"main.go/modules/user/model"
	"main.go/worker"
	"time"
)

func (biz *RegisterUserBiz) NewForgotPassword(ctx context.Context, data *model.ForgotPassword) error {
	if data.Email == "" {
		return common.ErrEmailRequire
	}
	_, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return common.ErrEmailNoExist(errors.New("user don't exist"))
	}

	//send email

	taskPayload := worker.PayloadSendEmailForgotPassword{
		Email: data.Email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendResetCodePassword),
		asynq.Unique(1 * time.Minute),
	}
	_ = biz.taskDistributor.DistributeTaskSendEmailForgotPassword(ctx, &taskPayload, opts...)

	return nil
}
