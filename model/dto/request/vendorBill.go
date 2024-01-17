package request

import "github.com/gin-gonic/gin"

type CreateVendorBill struct {
	Amount float64 `json:"amount"`
}

func (req *CreateVendorBill) FromGinContext(gc *gin.Context) error {
	if err := gc.BindJSON(req); err != nil {
		return err
	}

	return nil
}
