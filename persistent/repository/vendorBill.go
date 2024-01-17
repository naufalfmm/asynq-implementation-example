package repository

import (
	"asynq-implementation-example/model/dao"
	"context"
	"time"

	"github.com/google/uuid"
)

type VendorBillRepository interface {
	Create(ctx context.Context, createdDatum dao.VendorBill) (dao.VendorBill, error)
}

type vendorBillRepository struct{}

func NewVendorBillRepository() VendorBillRepository {
	return &vendorBillRepository{}
}

func (r *vendorBillRepository) Create(ctx context.Context, createdDatum dao.VendorBill) (dao.VendorBill, error) {
	createdDatum.ID = uuid.New().String()
	createdDatum.CreatedAt = time.Now()
	createdDatum.UpdatedAt = createdDatum.CreatedAt

	dao.VendorBills.Store(createdDatum.ID, createdDatum)
	return createdDatum, nil
}
