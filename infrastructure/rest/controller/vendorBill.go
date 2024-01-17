package controller

import (
	"asynq-implementation-example/handler"
	"asynq-implementation-example/model/dto/request"
	"asynq-implementation-example/model/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type vendorBillController struct {
	Handler handler.Handler
}

func NewVendorBillController(handl handler.Handler) vendorBillController {
	return vendorBillController{
		Handler: handl,
	}
}

func (c *vendorBillController) Create(gc *gin.Context) {
	var req request.CreateVendorBill
	if err := req.FromGinContext(gc); err != nil {
		response.NewJSONResponse(gc, http.StatusBadRequest, err.Error(), err)
		return
	}

	vb, err := c.Handler.VendorBillHandler.Create(gc.Request.Context(), req)
	if err != nil {
		response.NewJSONResponse(gc, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.NewJSONResponse(gc, http.StatusOK, "Success", response.NewVendorBillResponse(vb))
}
