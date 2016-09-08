// +build integration

package gopwa

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var flagSellerID = flag.String("seller-id", "", "Sandbox MWS Seller ID")
var flagAccessKeyID = flag.String("key-id", "", "Sandbox MWS Access Key ID")
var flagAccessKeySecret = flag.String("key-secret", "", "Sandbox MWS Access Key Secret")
var flagRegion = flag.Int("region", int(UK), "Region for accessing MWS (see: endpoints.go)")

var (
	setDetails                        = flag.String("set-order-reference-details", "", "AmazonOrderReferenceId for test: Test_SetOrderReferenceDetails")
	getsDetailsWithoutShippingAddress = flag.String("get-details-without-shipping-address", "", "AmazonOrderReferenceId for test: Test_GetOrderReferenceDetails_GetsDetailsWithoutShippingAddress")
	confirmOrderReference             = flag.String("confirm-order-reference", "", "AmazonOrderReferenceId for test: Test_ConfirmOrderReference")
	cancelOrderReference              = flag.String("cancel-order-reference", "", "AmazonOrderReferenceId for test: Test_CancelOrderReference")
	closeOrderReference               = flag.String("close-order-reference", "", "AmazonOrderReferenceId for test: Test_CloseOrderReference")
	authorize                         = flag.String("authorize", "", "AmazonOrderReferenceId for test: Test_Authorize")
	authorizeCapture                  = flag.String("authorize-capture", "", "AmazonOrderReferenceId for test: Test_Authorize_Capture")
	getAuthorizationDetails           = flag.String("get-authorization-details", "", "AmazonOrderReferenceId for test: Test_GetAuthorizationDetails")
	closeAuthorization                = flag.String("close-authorization", "", "AmazonOrderReferenceId for test: Test_CloseAuthorization")
	capture                           = flag.String("capture", "", "AmazonOrderReferenceId for test: Test_Capture")
	getCaptureDetails                 = flag.String("get-capture-details", "", "AmazonOrderReferenceId for test: Test_GetCaptureDetails")
	refund                            = flag.String("refund", "", "AmazonOrderReferenceId for test: Test_Refund")
	getRefundDetails                  = flag.String("get-refund-details", "", "AmazonOrderReferenceId for test: Test_GetRefundDetails")
)

func Test_GetServieStatus(t *testing.T) {
	if !requiredFlagsSet() {
		t.SkipNow()
	}
	pwa := setupPwa()

	resp, err := pwa.GetServiceStatus()
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Result.Status)
	assert.NotEmpty(t, resp.Metadata.RequestId)
	assert.NotZero(t, resp.Result.Timestamp)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -get-details-without-shipping-address={amazonOrderReferenceId}
func Test_GetOrderReferenceDetails_GetsDetailsWithoutShippingAddress(t *testing.T) {
	if !requiredFlagsSet(*getsDetailsWithoutShippingAddress) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *getsDetailsWithoutShippingAddress
	pwa := setupPwa()

	resp, err := pwa.GetOrderReferenceDetails(amazonOrderReferenceId, "")
	assert.NoError(t, err)

	assert.NotNil(t, resp.Result)
	assert.NotEmpty(t, resp.Metadata.RequestId)

	orderDetails := resp.Result.OrderReferenceDetails
	assert.Equal(t, amazonOrderReferenceId, orderDetails.AmazonOrderReferenceId)
	assert.Equal(t, "Sandbox", orderDetails.ReleaseEnvironment)
	assert.Equal(t, "Physical", orderDetails.Destination.DestinationType)
	assert.Equal(t, OrderRefStateDraft, orderDetails.OrderReferenceStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -set-order-reference-details={amazonOrderReferenceId}
func Test_SetOrderReferenceDetails(t *testing.T) {
	if !requiredFlagsSet(*setDetails) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *setDetails
	pwa := setupPwa()

	resp := setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	assert.NotNil(t, resp.Result)
	assert.NotEmpty(t, resp.Metadata.RequestId)

	orderDetails := resp.Result.OrderReferenceDetails
	assert.Equal(t, amazonOrderReferenceId, orderDetails.AmazonOrderReferenceId)
	assert.Equal(t, "Sandbox", orderDetails.ReleaseEnvironment)
	assert.Equal(t, "Physical", orderDetails.Destination.DestinationType)
	assert.Equal(t, OrderRefStateDraft, orderDetails.OrderReferenceStatus.State)
	assert.Equal(t, "TestSellerNote", orderDetails.SellerNote)
	assert.Equal(t, "GBP", orderDetails.OrderTotal.CurrencyCode)
	assert.Equal(t, "10.00", orderDetails.OrderTotal.Amount)
	assert.Equal(t, "abc123", orderDetails.SellerOrderAttributes.SellerOrderId)
	assert.Equal(t, "TestStore", orderDetails.SellerOrderAttributes.StoreName)
	assert.Equal(t, "TestCustomInformation", orderDetails.SellerOrderAttributes.CustomInformation)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -confirm-order-reference={amazonOrderReferenceId}
func Test_ConfirmOrderReference(t *testing.T) {
	if !requiredFlagsSet(*confirmOrderReference) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *confirmOrderReference
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	confirmOrderRefResp, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	assert.NotEmpty(t, confirmOrderRefResp.Metadata.RequestId)

	resp, err := pwa.GetOrderReferenceDetails(amazonOrderReferenceId, "")
	assert.NoError(t, err)

	orderDetails := resp.Result.OrderReferenceDetails
	assert.Equal(t, amazonOrderReferenceId, orderDetails.AmazonOrderReferenceId)
	assert.Equal(t, OrderRefStateOpen, orderDetails.OrderReferenceStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -cancel-order-reference={amazonOrderReferenceId}
func Test_CancelOrderReference(t *testing.T) {
	if !requiredFlagsSet(*cancelOrderReference) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *cancelOrderReference
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	cancelOrderRefResp, err := pwa.CancelOrderReference(amazonOrderReferenceId, "Test: _cancellation- ~reason~")
	assert.NoError(t, err)

	assert.NotEmpty(t, cancelOrderRefResp.Metadata.RequestId)

	resp, err := pwa.GetOrderReferenceDetails(amazonOrderReferenceId, "")
	assert.NoError(t, err)

	orderDetails := resp.Result.OrderReferenceDetails
	assert.Equal(t, amazonOrderReferenceId, orderDetails.AmazonOrderReferenceId)
	assert.Equal(t, OrderRefStateCanceled, orderDetails.OrderReferenceStatus.State)
	assert.Equal(t, "Test: _cancellation- ~reason~", orderDetails.OrderReferenceStatus.ReasonDescription)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -close-order-reference={amazonOrderReferenceId}
func Test_CloseOrderReference(t *testing.T) {
	if !requiredFlagsSet(*closeOrderReference) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *closeOrderReference
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	closeOrderRefResp, err := pwa.CloseOrderReference(amazonOrderReferenceId, "Test closure reason")
	assert.NoError(t, err)

	assert.NotEmpty(t, closeOrderRefResp.Metadata.RequestId)

	resp, err := pwa.GetOrderReferenceDetails(amazonOrderReferenceId, "")
	assert.NoError(t, err)

	orderDetails := resp.Result.OrderReferenceDetails
	assert.Equal(t, amazonOrderReferenceId, orderDetails.AmazonOrderReferenceId)
	assert.Equal(t, OrderRefStateClosed, orderDetails.OrderReferenceStatus.State)
	assert.Equal(t, "Test closure reason", orderDetails.OrderReferenceStatus.ReasonDescription)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -authorize={amazonOrderReferenceId}
func Test_Authorize(t *testing.T) {
	if !requiredFlagsSet(*authorize) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *authorize
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	resp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Metadata.RequestId)

	authorizationDetails := resp.Result.AuthorizationDetails
	assert.Equal(t, "GBP", authorizationDetails.AuthorizationAmount.CurrencyCode)
	assert.Equal(t, "10.00", authorizationDetails.AuthorizationAmount.Amount)
	assert.Equal(t, authorizationReferenceId, authorizationDetails.AuthorizationReferenceId)
	assert.Equal(t, "Test SellerAuthorizationNote", authorizationDetails.SellerAuthorizationNote)
	assert.Equal(t, AuthStateOpen, authorizationDetails.AuthorizationStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -authorize-capture={amazonOrderReferenceId}
func Test_Authorize_Capture(t *testing.T) {
	if !requiredFlagsSet(*authorizeCapture) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *authorizeCapture
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	resp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, true, "TSoftDescriptor")
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Metadata.RequestId)

	authorizationDetails := resp.Result.AuthorizationDetails
	assert.Equal(t, "GBP", authorizationDetails.AuthorizationAmount.CurrencyCode)
	assert.Equal(t, "10.00", authorizationDetails.AuthorizationAmount.Amount)
	assert.Equal(t, authorizationReferenceId, authorizationDetails.AuthorizationReferenceId)
	assert.Equal(t, "Test SellerAuthorizationNote", authorizationDetails.SellerAuthorizationNote)
	assert.Equal(t, AuthStateClosed, authorizationDetails.AuthorizationStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -get-authorization-details={amazonOrderReferenceId}
func Test_GetAuthorizationDetails(t *testing.T) {
	if !requiredFlagsSet(*getAuthorizationDetails) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *getAuthorizationDetails
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	authorizeResp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	resp, err := pwa.GetAuthorizationDetails(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId)
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Metadata.RequestId)

	authorizationDetails := resp.Result.AuthorizationDetails
	assert.Equal(t, authorizationReferenceId, authorizationDetails.AuthorizationReferenceId)
	assert.Equal(t, AuthStateOpen, authorizationDetails.AuthorizationStatus.State)
	assert.Equal(t, "Test SellerAuthorizationNote", authorizationDetails.SellerAuthorizationNote)
	assert.Equal(t, Price{CurrencyCode: "GBP", Amount: "10.00"}, authorizationDetails.AuthorizationAmount)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -close-authorization={amazonOrderReferenceId}
func Test_CloseAuthorization(t *testing.T) {
	if !requiredFlagsSet(*closeAuthorization) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *closeAuthorization
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	authorizeResp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	closeAuthorizationResp, err := pwa.CloseAuthorization(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId, "Test closure reason")
	assert.NoError(t, err)

	assert.NotEmpty(t, closeAuthorizationResp.Metadata.RequestId)

	resp, err := pwa.GetAuthorizationDetails(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId)
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Metadata.RequestId)

	authorizationDetails := resp.Result.AuthorizationDetails
	assert.Equal(t, authorizationReferenceId, authorizationDetails.AuthorizationReferenceId)
	assert.Equal(t, AuthStateClosed, authorizationDetails.AuthorizationStatus.State)
	assert.Equal(t, "Test closure reason", authorizationDetails.AuthorizationStatus.ReasonDescription)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -capture={amazonOrderReferenceId}
func Test_Capture(t *testing.T) {
	if !requiredFlagsSet(*capture) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *capture
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	authorizeResp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	captureReferenceId := fmt.Sprint(time.Now().UnixNano())
	resp, err := pwa.Capture(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId, captureReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerCaptureNote", "TSoftDescriptor")

	assert.NotEmpty(t, resp.Metadata.RequestId)

	captureDetails := resp.Result.CaptureDetails
	assert.Equal(t, Price{CurrencyCode: "GBP", Amount: "10.00"}, captureDetails.CaptureAmount)
	assert.Equal(t, captureReferenceId, captureDetails.CaptureReferenceId)
	assert.Equal(t, "Test SellerCaptureNote", captureDetails.SellerCaptureNote)
	assert.Equal(t, CaptureStateCompleted, captureDetails.CaptureStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -get-capture-details={amazonOrderReferenceId}
func Test_GetCaptureDetails(t *testing.T) {
	if !requiredFlagsSet(*getCaptureDetails) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *getCaptureDetails
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	authorizeResp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	captureReferenceId := fmt.Sprint(time.Now().UnixNano())
	captureResp, err := pwa.Capture(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId, captureReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerCaptureNote", "TSoftDescriptor")

	resp, err := pwa.GetCaptureDetails(captureResp.Result.CaptureDetails.AmazonCaptureId)
	assert.NoError(t, err)

	assert.NotEmpty(t, resp.Metadata.RequestId)

	captureDetails := resp.Result.CaptureDetails
	assert.Equal(t, Price{CurrencyCode: "GBP", Amount: "10.00"}, captureDetails.CaptureAmount)
	assert.Equal(t, captureReferenceId, captureDetails.CaptureReferenceId)
	assert.Equal(t, "Test SellerCaptureNote", captureDetails.SellerCaptureNote)
	assert.Equal(t, CaptureStateCompleted, captureDetails.CaptureStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -refund={amazonOrderReferenceId}
func Test_Refund(t *testing.T) {
	if !requiredFlagsSet(*refund) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *refund
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	authorizeResp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	captureReferenceId := fmt.Sprint(time.Now().UnixNano())
	captureResp, err := pwa.Capture(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId, captureReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerCaptureNote", "TSoftDescriptor")

	refundReferenceId := fmt.Sprint(time.Now().UnixNano())
	resp, err := pwa.Refund(captureResp.Result.CaptureDetails.AmazonCaptureId, refundReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerRefundNote", "TSoftDescriptor")

	assert.NotEmpty(t, resp.Metadata.RequestId)

	refundDetails := resp.Result.RefundDetails
	assert.Equal(t, Price{CurrencyCode: "GBP", Amount: "10.00"}, refundDetails.RefundAmount)
	assert.Equal(t, refundReferenceId, refundDetails.RefundReferenceId)
	assert.Equal(t, "Test SellerRefundNote", refundDetails.SellerRefundNote)
	assert.Equal(t, RefundStatePending, refundDetails.RefundStatus.State)
}

// Generate a new Amazon Order Reference ID and add the followng flag to the test:
//   -get-refund-details={amazonOrderReferenceId}
func Test_GetRefundDetails(t *testing.T) {
	if !requiredFlagsSet(*getRefundDetails) {
		t.SkipNow()
	}
	amazonOrderReferenceId := *getRefundDetails
	pwa := setupPwa()

	setOrderReferenceDetails(t, amazonOrderReferenceId, pwa)

	_, err := pwa.ConfirmOrderReference(amazonOrderReferenceId)
	assert.NoError(t, err)

	authorizationReferenceId := fmt.Sprint(time.Now().UnixNano())
	authorizeResp, err := pwa.Authorize(amazonOrderReferenceId, authorizationReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerAuthorizationNote", 0, false, "")
	assert.NoError(t, err)

	captureReferenceId := fmt.Sprint(time.Now().UnixNano())
	captureResp, err := pwa.Capture(authorizeResp.Result.AuthorizationDetails.AmazonAuthorizationId, captureReferenceId, Price{CurrencyCode: "GBP", Amount: "10.00"}, "Test SellerCaptureNote", "TSoftDescriptor")

	refundReferenceId := fmt.Sprint(time.Now().UnixNano())
	refundResp, err := pwa.Refund(captureResp.Result.CaptureDetails.AmazonCaptureId, refundReferenceId, Price{CurrencyCode: "GBP", Amount: "5.00"}, "Test SellerRefundNote", "TSoftDescriptor")

	resp, err := pwa.GetRefundDetails(refundResp.Result.RefundDetails.AmazonRefundId)

	assert.NotEmpty(t, resp.Metadata.RequestId)

	refundDetails := resp.Result.RefundDetails
	assert.Equal(t, Price{CurrencyCode: "GBP", Amount: "5.00"}, refundDetails.RefundAmount)
	assert.Equal(t, refundReferenceId, refundDetails.RefundReferenceId)
	assert.Equal(t, "Test SellerRefundNote", refundDetails.SellerRefundNote)
	assert.Equal(t, RefundStatePending, refundDetails.RefundStatus.State)
}

func setupPwa() *PayWithAmazon {
	return New(*flagSellerID, *flagAccessKeyID, *flagAccessKeySecret, Region(*flagRegion), Sandbox)
}

func setOrderReferenceDetails(t *testing.T, amazonOrderReferenceId string, pwa *PayWithAmazon) *SetOrderReferenceDetailsResponse {
	resp, err := pwa.SetOrderReferenceDetails(
		amazonOrderReferenceId,
		OrderReferenceAttributes{
			SellerNote: "TestSellerNote",
			OrderTotal: OrderTotal{
				CurrencyCode: "GBP",
				Amount:       "10.00",
			},
			SellerOrderAttributes: SellerOrderAttributes{
				SellerOrderId:     "abc123",
				StoreName:         "TestStore",
				CustomInformation: "TestCustomInformation",
			},
		},
	)
	assert.NoError(t, err)

	return resp
}

func requiredFlagsSet(requiredFlags ...string) bool {
	if *flagSellerID == "" ||
		*flagAccessKeyID == "" ||
		*flagAccessKeySecret == "" {
		return false
	}

	for _, requiredFlag := range requiredFlags {
		if requiredFlag == "" {
			return false
		}
	}

	return true
}
