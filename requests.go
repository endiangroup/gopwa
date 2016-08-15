package gopwa

import (
	"fmt"
	"net/url"
)

type Request interface {
	Action() string
	AddValues(url.Values) url.Values
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752010
type Authorize struct {
	AmazonOrderReferenceId   string
	AuthorizationReferenceId string
	AuthorizationAmount      Price
	SellerAuthorizationNote  string
	TransactionTimeout       uint
	CaptureNow               bool
	SoftDescriptor           string
}

func (req Authorize) Action() string {
	return "Authorize"
}

func (req Authorize) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)
	v.Set("AuthorizationReferenceId", req.AuthorizationReferenceId)
	v.Set("SellerAuthorizationNote", req.SellerAuthorizationNote)
	v.Set("SoftDescriptor", req.SoftDescriptor)
	v.Set("TransactionTimeout", fmt.Sprintf("%d", req.TransactionTimeout))
	if req.CaptureNow {
		v.Set("CaptureNow", "true")
	} else {
		v.Set("CaptureNow", "false")
	}

	return req.AuthorizationAmount.AddValues("AuthorizationAmount.", v)
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201751990
type CancelOrderReference struct {
	AmazonOrderReferenceId string
	CancelationReason      string
}

func (req CancelOrderReference) Action() string {
	return "GetOrderReferenceDetails"
}

func (req CancelOrderReference) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)
	v.Set("CancelationReason", req.CancelationReason)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752040
type Capture struct {
	AmazonAuthorizationId string
	CaptureReferenceId    string
	CaptureAmount         Price
	SellerCaptureNote     string
	SoftDescriptor        string
}

func (req Capture) Action() string {
	return "Capture"
}

func (req Capture) AddValues(v url.Values) url.Values {
	v.Set("AmazonAuthorizationId", req.AmazonAuthorizationId)
	v.Set("CaptureReferenceId", req.CaptureReferenceId)
	v.Set("SellerCaptureNote", req.SellerCaptureNote)
	v.Set("SoftDescriptor", req.SoftDescriptor)

	return req.CaptureAmount.AddValues("CaptureAmount.", v)
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752070
type CloseAuthorization struct {
	AmazonAuthorizationId string
	ClosureReason         string
}

func (req CloseAuthorization) Action() string {
	return "CloseAuthorization"
}

func (req CloseAuthorization) AddValues(v url.Values) url.Values {
	v.Set("AmazonAuthorizationId", req.AmazonAuthorizationId)
	v.Set("ClosureReason", req.ClosureReason)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752000
type CloseOrderReference struct {
	AmazonOrderReferenceId string
	ClosureReason          string
}

func (req CloseOrderReference) Action() string {
	return "CloseOrderReference"
}

func (req CloseOrderReference) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)
	v.Set("ClosureReason", req.ClosureReason)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201751980
type ConfirmOrderReference struct {
	AmazonOrderReferenceId string
}

func (req ConfirmOrderReference) Action() string {
	return "ConfirmOrderReference"
}

func (req ConfirmOrderReference) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201751670
type CreateOrderReferenceForId struct {
	Id                       string
	IdType                   string
	InheritShippingAddress   bool
	ConfirmNow               bool
	OrderReferenceAttributes OrderReferenceAttributes
}

func (req CreateOrderReferenceForId) Action() string {
	return "CreateOrderReferenceForId"
}

func (req CreateOrderReferenceForId) AddValues(v url.Values) url.Values {
	v.Set("Id", req.Id)
	v.Set("IdType", req.IdType)
	if req.InheritShippingAddress {
		v.Set("InheritShippingAddress", "true")
	} else {
		v.Set("InheritShippingAddress", "false")
	}
	if req.ConfirmNow {
		v.Set("ConfirmNow", "true")
		req.OrderReferenceAttributes.AddValues("OrderReferenceAttributes.", v)
	} else {
		v.Set("ConfirmNow", "false")
	}

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752030
type GetAuthorizationDetails struct {
	AmazonAuthorizationId string
}

func (req GetAuthorizationDetails) Action() string {
	return "GetAuthorizationDetails"
}

func (req GetAuthorizationDetails) AddValues(v url.Values) url.Values {
	v.Set("AmazonAuthorizationId", req.AmazonAuthorizationId)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752060
type GetCaptureDetails struct {
	AmazonCaptureId string
}

func (req GetCaptureDetails) Action() string {
	return "GetCaptureDetails"
}

func (req GetCaptureDetails) AddValues(v url.Values) url.Values {
	v.Set("AmazonCaptureId", req.AmazonCaptureId)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201751970
type GetOrderReferenceDetails struct {
	AmazonOrderReferenceId string
	AddressConsentToken    string
}

func (req GetOrderReferenceDetails) Action() string {
	return "GetOrderReferenceDetails"
}

func (req GetOrderReferenceDetails) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)
	v.Set("AddressConsentToken", req.AddressConsentToken)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752100
type GetRefundDetails struct {
	AmazonRefundId string
}

func (req GetRefundDetails) Action() string {
	return "GetRefundDetails"
}

func (req GetRefundDetails) AddValues(v url.Values) url.Values {
	v.Set("AmazonRefundId", req.AmazonRefundId)

	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752110
type GetServiceStatus struct{}

func (req GetServiceStatus) Action() string {
	return "GetServiceStatus"
}

func (req GetServiceStatus) AddValues(v url.Values) url.Values {
	return v
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201752080
type Refund struct {
	AmazonCaptureId   string
	RefundReferenceId string
	RefundAmount      Price
	SellerRefundNote  string
	SoftDescriptor    string
}

func (req Refund) Action() string {
	return "Refund"
}

func (req Refund) AddValues(v url.Values) url.Values {
	v.Set("AmazonCaptureId", req.AmazonCaptureId)
	v.Set("RefundReferenceId", req.RefundReferenceId)
	v.Set("SellerRefundNote", req.SellerRefundNote)
	v.Set("SoftDescriptor", req.SoftDescriptor)

	return req.RefundAmount.AddValues("RefundAmount.", v)
}

// See: https://payments.amazon.co.uk/developer/documentation/apireference/201751960
type SetOrderReferenceDetails struct {
	AmazonOrderReferenceId   string
	OrderReferenceAttributes OrderReferenceAttributes
}

func (req SetOrderReferenceDetails) Action() string {
	return "SetOrderReferenceDetails"
}

func (req SetOrderReferenceDetails) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)

	return req.OrderReferenceAttributes.AddValues("OrderReferenceAttributes.", v)
}
