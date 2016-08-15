package gopwa

import (
	"encoding/xml"
	"net/url"
	"time"
)

type OrderTotal struct {
	XMLName      xml.Name `xml:"OrderTotal"`
	CurrencyCode string
	Amount       string
}

func (o OrderTotal) AddValues(prefix string, v url.Values) url.Values {
	v.Set(prefix+"Amount", o.Amount)
	v.Set(prefix+"CurrencyCode", o.CurrencyCode)

	return v
}

type SellerOrderAttributes struct {
	XMLName           xml.Name `xml:"SellerOrderAttributes"`
	SellerOrderId     string
	StoreName         string
	CustomInformation string
}

func (s SellerOrderAttributes) AddValues(prefix string, v url.Values) url.Values {
	v.Set(prefix+"SellerOrderId", s.SellerOrderId)
	v.Set(prefix+"StoreName", s.StoreName)
	v.Set(prefix+"CustomInformation", s.CustomInformation)

	return v
}

type Address struct {
	Name          string
	AddressLine1  string
	AddressLine2  string
	AddressLine3  string
	City          string
	County        string
	District      string
	StateOrRegion string
	PostalCode    string
	CountryCode   string
	Phone         string
}

type Buyer struct {
	XMLName xml.Name `xml:"Buyer"`
	Name    string
	Email   string
	Phone   string
}

type Destination struct {
	XMLName             xml.Name `xml:"Destination"`
	DestinationType     string
	PhysicalDestination Address
}

type Constraint struct {
	XMLName      xml.Name `xml:"Constraint"`
	ConstraintID string
	Description  string
}

type OrderReferenceStatus struct {
	XMLName             xml.Name `xml:"OrderReferenceStatus"`
	State               string
	LastUpdateTimestamp time.Time
	ReasonCode          string
	ReasonDescription   string
}

type OrderReferenceAttributes struct {
	PlatformId string
	SellerNote string
	OrderTotal
	SellerOrderAttributes
}

func (o OrderReferenceAttributes) AddValues(prefix string, v url.Values) url.Values {
	v.Set(prefix+"PlatformId", o.PlatformId)
	v.Set(prefix+"SellerNote", o.SellerNote)

	v = o.OrderTotal.AddValues(prefix+"OrderTotal.", v)

	return o.SellerOrderAttributes.AddValues(prefix+"SellerOrderAttributes.", v)
}

type OrderReferenceDetails struct {
	XMLName                xml.Name `xml:"OrderReferenceDetails"`
	AmazonOrderReferenceId string
	SellerNote             string
	PlatformId             string
	ReleaseEnvironment     string
	OrderLanguage          string
	CreationTimestamp      time.Time
	ExpirationTimestamp    time.Time
	BillingAddress         Address
	Constraints            []Constraint
	IdList                 []string
	Buyer
	OrderTotal
	Destination
	OrderReferenceStatus
	SellerOrderAttributes
}
