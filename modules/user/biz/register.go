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

func (biz *RegisterUserBiz) NewRegister(ctx context.Context, data *model.Register, expiry int) error {

	if data.Email == "" {
		return errors.New("email is request")
	}
	if _, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email}); err == nil {
		return errors.New("user already exists")
	}
	data.Salt = common.GetSalt(50)
	data.IsEMail = false
	data.Password = biz.hash.Hash(data.Salt + data.Password)
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return errors.New("error create user")
	}

	//task

	taskPayload := worker.PayloadSendVerifyEmail{
		Email: data.Email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendVerifyEmail),
	}
	_ = biz.taskDistributor.DistributeTaskSendVerifyEmail(ctx, &taskPayload, opts...)
	return nil
}
