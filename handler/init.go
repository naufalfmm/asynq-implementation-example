package handler

type Handler struct {
	VendorBillHandler VendorBillHandler
}

func NewHandler(vendorBillHand VendorBillHandler) Handler {
	return Handler{
		VendorBillHandler: vendorBillHand,
	}
}
