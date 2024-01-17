package notification

import (
	"asynq-implementation-example/consts"
	notificationDto "asynq-implementation-example/model/dto/notification"
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

type VendorBillNotification interface {
	Send(ctx context.Context, opts []asynq.Option, req notificationDto.VendorBillNotification) error
	Receive(ctx context.Context, opts []asynq.Option, task *asynq.Task) (notificationDto.VendorBillNotification, error)
}

type vendorBillNotification struct {
	client *asynq.Client
	logger zerolog.Logger
}

func NewVendorBillNotification(client *asynq.Client, logger zerolog.Logger) VendorBillNotification {
	return &vendorBillNotification{
		client: client,
		logger: logger,
	}
}

func (n *vendorBillNotification) Send(ctx context.Context, opts []asynq.Option, req notificationDto.VendorBillNotification) error {
	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	task := asynq.NewTask(consts.TypeVendorBillNotificationSend, jsonPayload, opts...)
	info, err := n.client.EnqueueContext(ctx, task, opts...)
	if err != nil {
		n.logger.Error().
			Str("type", task.Type()).
			Any("data", req.Data).
			Err(err)
		return err
	}

	n.logger.Info().
		Str("type", task.Type()).
		Any("data", req.Data).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueued task")

	return nil
}

func (n *vendorBillNotification) Receive(ctx context.Context, opts []asynq.Option, task *asynq.Task) (notificationDto.VendorBillNotification, error) {
	var req notificationDto.VendorBillNotification
	if err := json.Unmarshal(task.Payload(), &req); err != nil {
		return notificationDto.VendorBillNotification{}, err
	}

	return req, nil
}
