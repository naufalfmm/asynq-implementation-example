package infrastructure

import (
	"asynq-implementation-example/infrastructure/asynq"
	"asynq-implementation-example/infrastructure/rest"

	"github.com/gin-gonic/gin"

	asynqLib "github.com/hibiken/asynq"
)

type Infrastructure struct {
	Rest  rest.Rest
	Asynq asynq.Asynq
}

func NewInfrastructure(res rest.Rest, asyn asynq.Asynq) Infrastructure {
	return Infrastructure{
		Rest:  res,
		Asynq: asyn,
	}
}

func (i *Infrastructure) Register(ge *gin.Engine, am *asynqLib.ServeMux) {
	i.Rest.Register(ge)
	i.Asynq.Register(am)
}
