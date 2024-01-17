package routes

import (
	"asynq-implementation-example/consts"
	"asynq-implementation-example/infrastructure/asynq/controller"

	"github.com/hibiken/asynq"
)

type Routes struct {
	Controller controller.Controller
}

func NewRoutes(cont controller.Controller) Routes {
	return Routes{
		Controller: cont,
	}
}

func (r *Routes) Register(am *asynq.ServeMux) {
	am.HandleFunc(consts.TypeVendorBillNotificationSend, r.Controller.VendorBillController.Process)
}
