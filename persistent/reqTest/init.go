package reqtest

type ReqTest struct {
	VendorBillReqTest VendorBillReqTest
}

func NewReqTest(vendorBillReTe VendorBillReqTest) ReqTest {
	return ReqTest{
		VendorBillReqTest: vendorBillReTe,
	}
}
