package app

import (
	"asynq-implementation-example/handler"
	"asynq-implementation-example/infrastructure"
	"asynq-implementation-example/infrastructure/asynq"
	asynqController "asynq-implementation-example/infrastructure/asynq/controller"
	asynqRoutes "asynq-implementation-example/infrastructure/asynq/routes"
	"asynq-implementation-example/infrastructure/rest"
	restController "asynq-implementation-example/infrastructure/rest/controller"
	restRoutes "asynq-implementation-example/infrastructure/rest/routes"
	"asynq-implementation-example/middleware"
	"asynq-implementation-example/persistent"
	"asynq-implementation-example/persistent/notification"
	"asynq-implementation-example/persistent/repository"
	reqtest "asynq-implementation-example/persistent/reqTest"
	"asynq-implementation-example/resource/config"
	"asynq-implementation-example/resource/logger"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	asynqLib "github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

type App struct {
	Ge     *gin.Engine
	aro    asynqLib.RedisClientOpt
	am     *asynqLib.ServeMux
	conf   *config.EnvConfig
	logger logger.Logger
}

func Init() App {
	ge := gin.New()
	am := asynqLib.NewServeMux()

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	asynqRedisOpt := asynqLib.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
	}
	asynqClient := asynqLib.NewClient(asynqRedisOpt)

	zlog := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logr := logger.NewLogger()

	vendorBillRepo := repository.NewVendorBillRepository()
	repo := repository.NewRepository(vendorBillRepo)

	vendorBillNotif := notification.NewVendorBillNotification(asynqClient, zlog)
	notif := notification.NewNotification(vendorBillNotif)

	vendorBillReqTe := reqtest.NewVendorBillReqTest(zlog)
	reqte := reqtest.NewReqTest(vendorBillReqTe)

	persist := persistent.NewPersistent(repo, notif, reqte)

	vendorBillHandl := handler.NewVendorBillHandler(persist, zlog)
	handl := handler.NewHandler(vendorBillHandl)

	middl := middleware.Init()

	restVendorBillCtl := restController.NewVendorBillController(handl)
	restCtl := restController.NewController(restVendorBillCtl)

	restRout := restRoutes.NewRouter(restCtl, middl)
	re := rest.Init(restRout)

	asynqVendorBillCtl := asynqController.NewVendorBillController(handl)
	asynqCtl := asynqController.NewController(asynqVendorBillCtl)

	asynqRout := asynqRoutes.NewRoutes(asynqCtl)
	as := asynq.Init(asynqRout)

	infr := infrastructure.NewInfrastructure(re, as)
	infr.Register(ge, am)

	return App{
		Ge:     ge,
		am:     am,
		aro:    asynqRedisOpt,
		conf:   conf,
		logger: logr,
	}
}

func (app App) Run() {
	asynqServer := asynqLib.NewServer(app.aro, asynqLib.Config{
		Queues: map[string]int{
			"critical": 25,
			"default":  5,
		},
		Logger: &app.logger,
	})

	if err := asynqServer.Start(app.am); err != nil {
		panic(err)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", app.conf.Port),
		Handler: app.Ge,
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sc

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctxShutDown); err != nil {
		panic(err)
	}

	asynqServer.Shutdown()
}
