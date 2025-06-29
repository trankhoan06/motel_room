package worker

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	emailSend "main.go/email"
	modelEmail "main.go/modules/email/model"
	modelFollow "main.go/modules/follower/model"
	modelMappingNotify "main.go/modules/mapping_notify/model"
	"main.go/modules/user/model"
)

const (
	QueueSendVerifyEmail       = "send_verify_email"
	QueueSendResetCodePassword = "send_verify_reset_code_password"
	QueueDefault               = "default"
)

type TaskProcessor interface {
	Start() error
	TaskProcessSendEmailForgotPassword(ctx context.Context, task *asynq.Task) error
	TaskProcessSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}
type AccountStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}
type AccountCase interface {
	NewCreateVerifyCodeEmail(ctx context.Context, email string, expire int) (*modelEmail.CreateVerifyAccount, error)
	NewResendCodeEmail(ctx context.Context, email string, expire int, Type *modelEmail.TypeCode) (*modelEmail.CreateVerifyAccount, error)
}
type FollowStorage interface {
	ListFollow(ctx context.Context, cond map[string]interface{}) (*[]modelFollow.Follower, error)
}
type MappingNotifyStorage interface {
	CreateMappingNotify(ctx context.Context, data *[]modelMappingNotify.CreateMappingNotify, size int) error
}
type RedisTaskProcessor struct {
	server         *asynq.Server
	accountStorage AccountStorage
	Mailer         emailSend.Sender
	accountCase    AccountCase
	followStorage  FollowStorage
	mappingNotify  MappingNotifyStorage
}

func NewRedisTaskProcessor(redisOpts *asynq.RedisClientOpt, accountStorage AccountStorage, Mailer emailSend.Sender, accountCase AccountCase,
	follow FollowStorage, mapping MappingNotifyStorage) *RedisTaskProcessor {
	logger := NewLogger()
	redis.SetLogger(logger)
	server := asynq.NewServer(
		redisOpts,
		asynq.Config{
			Queues: map[string]int{
				QueueSendVerifyEmail:       5,
				QueueSendResetCodePassword: 5,
				QueueDefault:               5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("task type", task.Type()).
					Bytes("payload", task.Payload()).
					Msg("error when process task")
			}),
			Logger: logger,
		})
	return &RedisTaskProcessor{
		server:         server,
		accountStorage: accountStorage,
		Mailer:         Mailer,
		accountCase:    accountCase,
		followStorage:  follow,
		mappingNotify:  mapping,
	}
}
func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskForgotPassword, processor.TaskProcessSendEmailForgotPassword)
	mux.HandleFunc(TaskSendVerifyEmail, processor.TaskProcessSendVerifyEmail)
	mux.HandleFunc(TaskCreateNotify, processor.TaskProcessCreateNotify)

	return processor.server.Start(mux)

}
