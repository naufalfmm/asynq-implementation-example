package response

import (
	"asynq-implementation-example/model/dao"
	"time"
)

type VendorBillResponse struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVendorBillResponse(vb dao.VendorBill) VendorBillResponse {
	return VendorBillResponse{
		ID:        vb.ID,
		Amount:    vb.Amount,
		CreatedAt: vb.CreatedAt,
		UpdatedAt: vb.UpdatedAt,
	}
}
