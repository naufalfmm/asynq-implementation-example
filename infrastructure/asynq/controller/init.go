package controller

type Controller struct {
	VendorBillController vendorBillController
}

func NewController(vendorBillCont vendorBillController) Controller {
	return Controller{
		VendorBillController: vendorBillCont,
	}
}
