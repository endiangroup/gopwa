package gopwa

import "net/url"

type Request interface {
	Action() string
	AddValues(url.Values) url.Values
}

type SetOrderReferenceDetailsRequest struct {
	AmazonOrderReferenceId string
	OrderReferenceAttributes
}

func (req SetOrderReferenceDetailsRequest) Action() string {
	return "SetOrderReferenceDetails"
}

func (req SetOrderReferenceDetailsRequest) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)

	return req.OrderReferenceAttributes.AddValues("OrderReferenceAttributes.", v)
}

type GetOrderReferenceDetailsRequest struct {
	AmazonOrderReferenceId string
	AddressConsentToken    string
}

func (req GetOrderReferenceDetailsRequest) Action() string {
	return "GetOrderReferenceDetails"
}

func (req GetOrderReferenceDetailsRequest) AddValues(v url.Values) url.Values {
	v.Set("AmazonOrderReferenceId", req.AmazonOrderReferenceId)
	v.Set("AddressConsentToken", req.AddressConsentToken)

	return v
}
