package gopwa

import "time"

type AuthorizeResponse struct {
	AuthorizeResult
	ResponseMetadata
}

type AuthorizeResult struct {
	AuthorizationDetails AuthorizationDetails
}

type CancelOrderReferenceResponse struct {
	ResponseMetadata
}

type CaptureResponse struct {
	CaptureResult
	ResponseMetadata
}

type CaptureResult struct {
	CaptureDetails CaptureDetails
}

type CloseAuthorizationResponse struct {
	ResponseMetadata
}

type CloseOrderReferenceResponse struct {
	ResponseMetadata
}

type ConfirmOrderReferenceResponse struct {
	ResponseMetadata
}

type CreateOrderReferenceForIdResponse struct {
	CreateOrderReferenceForIdResult
	ResponseMetadata
}

type CreateOrderReferenceForIdResult struct {
	OrderReferenceDetails OrderReferenceDetails
}

type GetAuthorizationDetailsResponse struct {
	GetAuthorizationDetailsResult
	ResponseMetadata
}

type GetAuthorizationDetailsResult struct {
	AuthorizationDetails
}

type GetCaptureDetailsResponse struct {
	GetCaptureDetailsResult
	ResponseMetadata
}

type GetCaptureDetailsResult struct {
	CaptureDetails CaptureDetails
}

type GetOrderReferenceDetailsResponse struct {
	GetOrderReferenceDetailsResult
	ResponseMetadata
}

type GetOrderReferenceDetailsResult struct {
	OrderReferenceDetails OrderReferenceDetails
}

type GetRefundDetailsResponse struct {
	GetRefundDetailsResult
	ResponseMetadata
}

type GetRefundDetailsResult struct {
	RefundDetails RefundDetails
}

type GetServiceStatusResponse struct {
	GetServiceStatusResult
	ResponseMetadata
}

type GetServiceStatusResult struct {
	Status    string
	Timestamp time.Time
	MessageId string
	Messages  []Message
}

type RefundResponse struct {
	RefundResult
	ResponseMetadata
}

type RefundResult struct {
	RefundDetails RefundDetails
}

type SetOrderReferenceDetailsResponse struct {
	SetOrderReferenceDetailsResult
	ResponseMetadata
}

type SetOrderReferenceDetailsResult struct {
	OrderReferenceDetails OrderReferenceDetails
}
