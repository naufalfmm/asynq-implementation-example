package reqtest

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
)

type VendorBillReqTest interface {
	PostEmpty(ctx context.Context) error
}

type vendorBillReqTest struct {
	logger zerolog.Logger
}

func NewVendorBillReqTest(lgr zerolog.Logger) VendorBillReqTest {
	return &vendorBillReqTest{
		logger: lgr,
	}
}

func (v *vendorBillReqTest) PostEmpty(ctx context.Context) error {
	var client = &http.Client{}

	request, err := http.NewRequest("POST", "http://localhost:9000/", nil)
	if err != nil {
		v.logger.Error().Err(err).Msg("Post Empty Error")
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		v.logger.Error().Err(err).Msg("Post Empty Error")
		return err
	}
	defer response.Body.Close()

	v.logger.Info().Msg("Post Empty Success")

	return nil
}
