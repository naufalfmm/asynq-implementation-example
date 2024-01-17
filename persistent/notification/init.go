package notification

type Notification struct {
	VendorBillNotification VendorBillNotification
}

func NewNotification(vendorBillNotif VendorBillNotification) Notification {
	return Notification{
		VendorBillNotification: vendorBillNotif,
	}
}
