package rest

import (
	"asynq-implementation-example/infrastructure/rest/routes"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	Routes routes.Routes
}

func Init(rout routes.Routes) Rest {
	return Rest{
		Routes: rout,
	}
}

func (r *Rest) Register(ge *gin.Engine) {
	r.Routes.Register(ge)
}

// func (r *Rest) Start(ge *gin.Engine) {
// 	ge.RedirectFixedPath = true

// 	r.server = &http.Server{
// 		Addr: ":8080",
// 		Handler: ge,
// 	}

// 	ge.Run()
// }

// func Start() Rest {
// 	ge := gin.New()
// 	ge.RedirectFixedPath = true

// 	vendorBillRepo := repository.NewVendorBillRepository()
// 	repo := repository.NewRepository(vendorBillRepo)

// 	asynqClient := asynq.NewClient(asynq.RedisClientOpt{
// 		Addr:     "localhost:6379",
// 		Password: "1234567890",
// 	})
// 	logger := zerolog.New(os.Stderr)

// 	persist := persistent.NewPersistent(repo, notification.Notification{})

// 	vendorBillHandl := handler.NewVendorBillHandler(persist, logger)

// 	vendorBillCont := controller.NewVendorBillController(vendorBillHandl)
// 	cont := controller.NewController(vendorBillCont)

// 	rout := router.NewRouter(ge, cont)
// 	rout.Init()

// 	ht := Rest{
// 		server: &http.Server{
// 			Addr:    ":8080",
// 			Handler: ge,
// 		},
// 	}

// 	go func() {
// 		if err := ht.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
// 			logger.Error().Err(err)
// 		}
// 	}()

// 	return ht
// }
