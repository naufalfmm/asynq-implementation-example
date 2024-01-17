package controller

import (
	"asynq-implementation-example/handler"
	"context"

	"github.com/hibiken/asynq"
)

type vendorBillController struct {
	Handler handler.Handler
}

func NewVendorBillController(handl handler.Handler) vendorBillController {
	return vendorBillController{
		Handler: handl,
	}
}

func (c *vendorBillController) Process(ctx context.Context, task *asynq.Task) error {
	if err := c.Handler.VendorBillHandler.ProcessNotification(ctx, task); err != nil {
		return err
	}

	return nil
}
