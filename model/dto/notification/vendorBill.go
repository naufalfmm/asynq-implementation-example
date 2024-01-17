package notificationDto

type CreatedVendorBill struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type VendorBillNotificationType string

const (
	VendorBillGeneratedVendorBillNotificationType VendorBillNotificationType = "vendor-bill-generated"
)

type VendorBillNotification struct {
	Type VendorBillNotificationType `json:"type"`
	Data interface{}                `json:"data"`
}
