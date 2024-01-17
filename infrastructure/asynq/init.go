package asynq

import (
	"asynq-implementation-example/infrastructure/asynq/routes"

	"github.com/hibiken/asynq"
)

type Asynq struct {
	Routes routes.Routes
}

func Init(rout routes.Routes) Asynq {
	return Asynq{
		Routes: rout,
	}
}

func (r *Asynq) Register(am *asynq.ServeMux) {
	r.Routes.Register(am)
}

// func Start() {
// 	vendorBillRepo := repository.NewVendorBillRepository()
// 	repo := repository.NewRepository(vendorBillRepo)

// 	asynqClient := asynq.NewClient(asynq.RedisClientOpt{
// 		Addr:     "localhost:6379",
// 		Password: "1234567890",
// 	})
// 	logger := zerolog.New(os.Stderr)

// 	vendorBillNotif := notification.NewVendorBillNotification(asynqClient, logger)
// 	notif := notification.NewNotification(vendorBillNotif)

// 	persist := persistent.NewPersistent(repo, notif)

// 	vendorBillHandl := handler.NewVendorBillHandler(persist, logger)

// 	vendorBillCont := controller.NewVendorBillController(vendorBillHandl)
// 	cont := controller.NewController(vendorBillCont)

// 	mux := asynq.NewServeMux()

// 	rout := router.NewRouter(mux, cont)
// 	rout.Init()

// 	asynq.NewServer(asynq.RedisClientOpt{Addr: "localhost:6379", Password: "1234567890"},
// 		asynq.Config{
// 			Queues: map[string]int{
// 				"critical": 25,
// 				"default":  5,
// 			},
// 			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
// 				log.Error().Err(err).Str("type", task.Type()).
// 					Bytes("payload", task.Payload()).Msg("process task failed")
// 			}),
// 		}).Start(mux)
// }
