package gopwa

import "time"

type AuthorizeResponse struct {
	Result   AuthorizeResult  `xml:"AuthorizeResult"`
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type AuthorizeResult struct {
	AuthorizationDetails `xml:"AuthorizationDetails"`
}

type CancelOrderReferenceResponse struct {
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type CaptureResponse struct {
	Result   CaptureResult    `xml:"CaptureResult"`
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type CaptureResult struct {
	CaptureDetails `xml:"CaptureDetails"`
}

type CloseAuthorizationResponse struct {
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type CloseOrderReferenceResponse struct {
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type ConfirmOrderReferenceResponse struct {
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type CreateOrderReferenceForIdResponse struct {
	Result   CreateOrderReferenceForIdResult `xml:"CreateOrderReferenceForIdResult"`
	Metadata ResponseMetadata                `xml:"ResponseMetadata"`
}

type CreateOrderReferenceForIdResult struct {
	OrderReferenceDetails `xml:"OrderReferenceDetails"`
}

type GetAuthorizationDetailsResponse struct {
	Result   GetAuthorizationDetailsResult `xml:"GetAuthorizationDetailsResult"`
	Metadata ResponseMetadata              `xml:"ResponseMetadata"`
}

type GetAuthorizationDetailsResult struct {
	AuthorizationDetails `xml:"AuthorizationDetails"`
}

type GetCaptureDetailsResponse struct {
	Result   GetCaptureDetailsResult `xml:"GetCaptureDetailsResult"`
	Metadata ResponseMetadata        `xml:"ResponseMetadata"`
}

type GetCaptureDetailsResult struct {
	CaptureDetails `xml:"CaptureDetails"`
}

type GetOrderReferenceDetailsResponse struct {
	Result   GetOrderReferenceDetailsResult `xml:"GetOrderReferenceDetailsResult"`
	Metadata ResponseMetadata               `xml:"ResponseMetadata"`
}

type GetOrderReferenceDetailsResult struct {
	OrderReferenceDetails `xml:"OrderReferenceDetails"`
}

type GetRefundDetailsResponse struct {
	Result   GetRefundDetailsResult `xml:"GetRefundDetailsResult"`
	Metadata ResponseMetadata       `xml:"ResponseMetadata"`
}

type GetRefundDetailsResult struct {
	RefundDetails `xml:"RefundDetails"`
}

type GetServiceStatusResponse struct {
	Result   GetServiceStatusResult `xml:"GetServiceStatusResult"`
	Metadata ResponseMetadata       `xml:"ResponseMetadata"`
}

type GetServiceStatusResult struct {
	Status    string
	Timestamp time.Time
	MessageId string
	Messages  []Message
}

type RefundResponse struct {
	Result   RefundResult     `xml:"RefundResult"`
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type RefundResult struct {
	RefundDetails `xml:"RefundDetails"`
}

type SetOrderReferenceDetailsResponse struct {
	Result   SetOrderReferenceDetailsResult `xml:"SetOrderReferenceDetailsResult"`
	Metadata ResponseMetadata               `xml:"ResponseMetadata"`
}

type SetOrderReferenceDetailsResult struct {
	OrderReferenceDetails `xml:"OrderReferenceDetails"`
}
