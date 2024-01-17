package repository

type Repository struct {
	VendorBillRepository VendorBillRepository
}

func NewRepository(vendorBillRepo VendorBillRepository) Repository {
	return Repository{
		VendorBillRepository: vendorBillRepo,
	}
}
