package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	modelMappingNotify "main.go/modules/mapping_notify/model"
)

const TaskCreateNotify = "task:create_notify"

type PayloadCreateNotify struct {
	UserId         int `json:"user_id"`
	NotificationId int `json:"notification_id"`
}

func (dis *RedisTaskDistributor) DistributeTaskCreateNotify(
	ctx context.Context,
	payload *PayloadCreateNotify,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}
	task := asynq.NewTask(TaskCreateNotify, jsonPayload, opts...)
	info, err := dis.Client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}
func (pro *RedisTaskProcessor) TaskProcessCreateNotify(ctx context.Context, task *asynq.Task) error {
	var payload PayloadCreateNotify
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	users, err := pro.followStorage.ListFollow(ctx, map[string]interface{}{"user_id": payload.UserId})
	if err != nil {
		return fmt.Errorf("error user: %w", err)
	}
	batch := []modelMappingNotify.CreateMappingNotify{}
	for _, value := range *users {
		batch = append(batch, modelMappingNotify.CreateMappingNotify{
			UserId:         value.UserId,
			NotificationId: payload.NotificationId,
		})
	}
	if err := pro.mappingNotify.CreateMappingNotify(ctx, &batch, 1000); err != nil {
		return fmt.Errorf("error when create mapping notify: %w", err)
	}

	log.Info().Msg("send verify email success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("processed task")
	return nil
}
