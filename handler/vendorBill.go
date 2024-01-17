package handler

import (
	"asynq-implementation-example/model/dao"
	"asynq-implementation-example/model/dto/request"
	"asynq-implementation-example/persistent"
	"context"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"

	notificationDto "asynq-implementation-example/model/dto/notification"
)

type VendorBillHandler interface {
	Create(ctx context.Context, req request.CreateVendorBill) (dao.VendorBill, error)
	ProcessNotification(ctx context.Context, task *asynq.Task) error
}

type vendorBillHandler struct {
	persistent persistent.Persistent
	logger     zerolog.Logger
}

func NewVendorBillHandler(persist persistent.Persistent, logger zerolog.Logger) VendorBillHandler {
	return &vendorBillHandler{
		persistent: persist,
		logger:     logger,
	}
}

func (h *vendorBillHandler) Create(ctx context.Context, req request.CreateVendorBill) (dao.VendorBill, error) {
	createdVendorBill, err := h.persistent.Repository.VendorBillRepository.Create(ctx, dao.VendorBill{
		Amount: req.Amount,
	})
	if err != nil {
		return dao.VendorBill{}, err
	}

	if err := h.persistent.Notification.VendorBillNotification.Send(ctx, nil, notificationDto.VendorBillNotification{
		Type: notificationDto.VendorBillGeneratedVendorBillNotificationType,
		Data: notificationDto.CreatedVendorBill{
			ID:     createdVendorBill.ID,
			Amount: createdVendorBill.Amount,
		},
	}); err != nil {
		return dao.VendorBill{}, err
	}

	return createdVendorBill, nil
}

func (h *vendorBillHandler) ProcessNotification(ctx context.Context, task *asynq.Task) error {
	req, err := h.persistent.Notification.VendorBillNotification.Receive(ctx, nil, task)
	if err != nil {
		return err
	}

	if err := h.persistent.ReqTest.VendorBillReqTest.PostEmpty(ctx); err != nil {
		return err
	}

	h.logger.Info().Str("type", string(req.Type)).Interface("data", req.Data).Msg("Notification Processed")

	return nil
}
