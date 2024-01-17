package routes

import (
	"asynq-implementation-example/consts"
	"asynq-implementation-example/infrastructure/rest/controller"
	"asynq-implementation-example/middleware"
	"asynq-implementation-example/model/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Controller controller.Controller
	Middleware middleware.Middleware
}

func NewRouter(cont controller.Controller, middl middleware.Middleware) Routes {
	return Routes{
		Controller: cont,
		Middleware: middl,
	}
}

func (r *Routes) Register(ge *gin.Engine) {
	ge.Use(r.Middleware.ImplementCors())
	ge.RedirectFixedPath = true

	ge.GET("/", func(gc *gin.Context) {
		response.NewJSONResponse(gc, http.StatusOK, "Hello World!", nil)
	})

	ge.NoRoute(func(gc *gin.Context) {
		response.NewJSONResponse(gc, http.StatusNotFound, "", consts.ErrPathNotFound)
	})

	vendorBill := ge.Group("/vendor-bills")
	vendorBill.POST("", r.Controller.VendorBillController.Create)
}
