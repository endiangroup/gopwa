package gopwa

type AuthorizeResponse struct {
	Result   *AuthorizeResult `xml:"AuthorizeResult"`
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type AuthorizeResult struct {
	AuthorizationDetails AuthorizationDetails
}

type CancelOrderReferenceResponse struct {
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type CaptureResponse struct {
	Result   *CaptureResult   `xml:"CaptureResult"`
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type CaptureResult struct {
	CaptureDetails CaptureDetails
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
	Result   *CreateOrderReferenceForIdResult `xml:"CreateOrderReferenceForIdResult"`
	Metadata ResponseMetadata                 `xml:"ResponseMetadata"`
}

type CreateOrderReferenceForIdResult struct {
	OrderReferenceDetails OrderReferenceDetails
}

type GetAuthorizationDetailsResponse struct {
	Result   *GetAuthorizationDetailsResult `xml:"GetAuthorizationDetailsResult"`
	Metadata ResponseMetadata               `xml:"ResponseMetadata"`
}

type GetAuthorizationDetailsResult struct {
	AuthorizationDetails AuthorizationDetails
}

type GetCaptureDetailsResponse struct {
	Result   *GetCaptureDetailsResult `xml:"GetCaptureDetailsResult"`
	Metadata ResponseMetadata         `xml:"ResponseMetadata"`
}

type GetCaptureDetailsResult struct {
	CaptureDetails CaptureDetails
}

type GetMerchantAccountStatusResponse struct {
	Result   *GetMerchantAccountStatusResult `xml:"GetMerchantAccountStatusResult"`
	Metadata ResponseMetadata                `xml:"ResponseMetadata"`
}

type GetMerchantAccountStatusResult struct {
	AccountStatus string
}

type GetOrderReferenceDetailsResponse struct {
	Result   *GetOrderReferenceDetailsResult `xml:"GetOrderReferenceDetailsResult"`
	Metadata ResponseMetadata                `xml:"ResponseMetadata"`
}

type GetOrderReferenceDetailsResult struct {
	OrderReferenceDetails OrderReferenceDetails
}

type GetRefundDetailsResponse struct {
	Result   *GetRefundDetailsResult `xml:"GetRefundDetailsResult"`
	Metadata ResponseMetadata        `xml:"ResponseMetadata"`
}

type GetRefundDetailsResult struct {
	RefundDetails RefundDetails
}

type GetServiceStatusResponse struct {
	Result   *GetServiceStatusResult `xml:"GetServiceStatusResult"`
	Metadata ResponseMetadata        `xml:"ResponseMetadata"`
}

type RefundResponse struct {
	Result   *RefundResult    `xml:"RefundResult"`
	Metadata ResponseMetadata `xml:"ResponseMetadata"`
}

type RefundResult struct {
	RefundDetails RefundDetails
}

type SetOrderReferenceDetailsResponse struct {
	Result   *SetOrderReferenceDetailsResult `xml:"SetOrderReferenceDetailsResult"`
	Metadata ResponseMetadata                `xml:"ResponseMetadata"`
}

type SetOrderReferenceDetailsResult struct {
	OrderReferenceDetails OrderReferenceDetails
}
