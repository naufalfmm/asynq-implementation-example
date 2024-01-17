package dao

import (
	"sync"
	"time"
)

var VendorBills sync.Map

type VendorBill struct {
	ID        string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
