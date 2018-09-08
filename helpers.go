package gopwa

func (pwa PayWithAmazon) Authorize(
	amazonOrderReferenceId string,
	authorizationReferenceId string,
	authorizationAmount Price,
	sellerAuthorizationNote string,
	transactionTimeout uint,
	captureNow bool,
	softDescriptor string,
) (*AuthorizeResponse, error) {
	result := &AuthorizeResponse{}
	request := Authorize{
		AmazonOrderReferenceId:   amazonOrderReferenceId,
		AuthorizationReferenceId: authorizationReferenceId,
		AuthorizationAmount:      authorizationAmount,
		SellerAuthorizationNote:  sellerAuthorizationNote,
		TransactionTimeout:       transactionTimeout,
		CaptureNow:               captureNow,
		SoftDescriptor:           softDescriptor,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) CancelOrderReference(amazonOrderReferenceId, cancelationReason string) (*CancelOrderReferenceResponse, error) {
	result := &CancelOrderReferenceResponse{}
	request := CancelOrderReference{
		AmazonOrderReferenceId: amazonOrderReferenceId,
		CancelationReason:      cancelationReason,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) Capture(
	amazonAuthorizationId string,
	captureReferenceId string,
	captureAmount Price,
	sellerCaptureNote string,
	softDescriptor string,
) (*CaptureResponse, error) {
	result := &CaptureResponse{}
	request := Capture{
		AmazonAuthorizationId: amazonAuthorizationId,
		CaptureReferenceId:    captureReferenceId,
		CaptureAmount:         captureAmount,
		SellerCaptureNote:     sellerCaptureNote,
		SoftDescriptor:        softDescriptor,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) CloseAuthorization(amazonAuthorizationId, closureReason string) (*CloseAuthorizationResponse, error) {
	result := &CloseAuthorizationResponse{}
	request := CloseAuthorization{
		AmazonAuthorizationId: amazonAuthorizationId,
		ClosureReason:         closureReason,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) CloseOrderReference(amazonOrderReferenceId, closureReason string) (*CloseOrderReferenceResponse, error) {
	result := &CloseOrderReferenceResponse{}
	request := CloseOrderReference{
		AmazonOrderReferenceId: amazonOrderReferenceId,
		ClosureReason:          closureReason,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) ConfirmOrderReference(amazonOrderReferenceId string) (*ConfirmOrderReferenceResponse, error) {
	result := &ConfirmOrderReferenceResponse{}
	request := ConfirmOrderReference{
		AmazonOrderReferenceId: amazonOrderReferenceId,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) CreateOrderReferenceForId(
	id string,
	idType string,
	inheritShippingAddress bool,
	confirmNow bool,
	orderReferenceAttributes OrderReferenceAttributes,
) (*CreateOrderReferenceForIdResponse, error) {
	result := &CreateOrderReferenceForIdResponse{}
	request := CreateOrderReferenceForId{
		Id:                       id,
		IdType:                   idType,
		InheritShippingAddress:   inheritShippingAddress,
		ConfirmNow:               confirmNow,
		OrderReferenceAttributes: orderReferenceAttributes,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) GetAuthorizationDetails(amazonAuthorizationId string) (*GetAuthorizationDetailsResponse, error) {
	result := &GetAuthorizationDetailsResponse{}
	request := GetAuthorizationDetails{
		AmazonAuthorizationId: amazonAuthorizationId,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) GetCaptureDetails(amazonCaptureId string) (*GetCaptureDetailsResponse, error) {
	result := &GetCaptureDetailsResponse{}
	request := GetCaptureDetails{
		AmazonCaptureId: amazonCaptureId,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) GetMerchantAccountStatus(sellerId string) (*GetMerchantAccountStatusResponse, error) {
	result := &GetMerchantAccountStatusResponse{}
	request := GetMerchantAccountStatus{
		SellerId: sellerId,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) GetOrderReferenceDetails(amazonOrderReferenceId, addressConsentToken string) (*GetOrderReferenceDetailsResponse, error) {
	result := &GetOrderReferenceDetailsResponse{}
	request := GetOrderReferenceDetails{
		AmazonOrderReferenceId: amazonOrderReferenceId,
		AddressConsentToken:    addressConsentToken,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) GetRefundDetails(amazonRefundId string) (*GetRefundDetailsResponse, error) {
	result := &GetRefundDetailsResponse{}
	request := GetRefundDetails{
		AmazonRefundId: amazonRefundId,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) GetServiceStatus() (*GetServiceStatusResponse, error) {
	result := &GetServiceStatusResponse{}

	if err := pwa.Do(GetServiceStatus{}, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) Refund(
	amazonCaptureId string,
	refundReferenceId string,
	refundAmount Price,
	sellerRefundNote string,
	softDescriptor string,
) (*RefundResponse, error) {
	result := &RefundResponse{}
	request := Refund{
		AmazonCaptureId:   amazonCaptureId,
		RefundReferenceId: refundReferenceId,
		RefundAmount:      refundAmount,
		SellerRefundNote:  sellerRefundNote,
		SoftDescriptor:    softDescriptor,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}

func (pwa PayWithAmazon) SetOrderReferenceDetails(amazonOrderReferenceId string, orderReferenceAttributes OrderReferenceAttributes) (*SetOrderReferenceDetailsResponse, error) {
	result := &SetOrderReferenceDetailsResponse{}
	request := SetOrderReferenceDetails{
		AmazonOrderReferenceId:   amazonOrderReferenceId,
		OrderReferenceAttributes: orderReferenceAttributes,
	}

	if err := pwa.Do(request, result); err != nil {
		return result, err
	}

	return result, nil
}
