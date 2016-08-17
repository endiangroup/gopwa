package gopwa

import "time"

type AuthorizeResponse struct {
	AuthorizeResult  AuthorizeResult
	ResponseMetadata ResponseMetadata
}

type AuthorizeResult struct {
	AuthorizationDetails AuthorizationDetails
}

type CancelOrderReferenceResponse struct {
	ResponseMetadata ResponseMetadata
}

type CaptureResponse struct {
	CaptureResult    CaptureResult
	ResponseMetadata ResponseMetadata
}

type CaptureResult struct {
	CaptureDetails CaptureDetails
}

type CloseAuthorizationResponse struct {
	ResponseMetadata ResponseMetadata
}

type CloseOrderReferenceResponse struct {
	ResponseMetadata ResponseMetadata
}

type ConfirmOrderReferenceResponse struct {
	ResponseMetadata ResponseMetadata
}

type CreateOrderReferenceForIdResponse struct {
	CreateOrderReferenceForIdResult CreateOrderReferenceForIdResult
	ResponseMetadata                ResponseMetadata
}

type CreateOrderReferenceForIdResult struct {
	OrderReferenceDetails OrderReferenceDetails
}

type GetAuthorizationDetailsResponse struct {
	GetAuthorizationDetailsResult GetAuthorizationDetailsResult
	ResponseMetadata              ResponseMetadata
}

type GetAuthorizationDetailsResult struct {
	AuthorizationDetails AuthorizationDetails
}

type GetCaptureDetailsResponse struct {
	GetCaptureDetailsResult GetCaptureDetailsResult
	ResponseMetadata        ResponseMetadata
}

type GetCaptureDetailsResult struct {
	CaptureDetails CaptureDetails
}

type GetOrderReferenceDetailsResponse struct {
	GetOrderReferenceDetailsResult GetOrderReferenceDetailsResult
	ResponseMetadata               ResponseMetadata
}

type GetOrderReferenceDetailsResult struct {
	OrderReferenceDetails OrderReferenceDetails
}

type GetRefundDetailsResponse struct {
	GetRefundDetailsResult GetRefundDetailsResult
	ResponseMetadata       ResponseMetadata
}

type GetRefundDetailsResult struct {
	RefundDetails RefundDetails
}

type GetServiceStatusResponse struct {
	GetServiceStatusResult GetServiceStatusResult
	ResponseMetadata       ResponseMetadata
}

type GetServiceStatusResult struct {
	Status    string
	Timestamp time.Time
	MessageId string
	Messages  []Message
}

type RefundResponse struct {
	RefundResult     RefundResult
	ResponseMetadata ResponseMetadata
}

type RefundResult struct {
	RefundDetails RefundDetails
}

type SetOrderReferenceDetailsResponse struct {
	SetOrderReferenceDetailsResult SetOrderReferenceDetailsResult
	ResponseMetadata               ResponseMetadata
}

type SetOrderReferenceDetailsResult struct {
	OrderReferenceDetails OrderReferenceDetails
}
