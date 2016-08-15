package gopwa

import "encoding/xml"

type SetOrderReferenceDetailsResponse struct {
	XMLName xml.Name `xml:"SetOrderReferenceDetailsResponse"`
	SetOrderReferenceDetailsResult
}

type SetOrderReferenceDetailsResult struct {
	XMLName xml.Name `xml:"SetOrderReferenceDetailsResult"`
	OrderReferenceDetails
}

type GetOrderReferenceDetailsResponse struct {
	XMLName xml.Name `xml:"GetOrderReferenceDetailsResponse"`
	GetOrderReferenceDetailsResult
}

type GetOrderReferenceDetailsResult struct {
	XMLName xml.Name `xml:"GetOrderReferenceDetailsResult"`
	OrderReferenceDetails
}
