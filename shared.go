package gopwa

import (
	"net/url"
	"time"
)

type ResponseMetadata struct {
	RequestId string
}

type OrderTotal struct {
	CurrencyCode string
	Amount       string
}

func (o OrderTotal) AddValues(prefix string, v url.Values) url.Values {
	if o.Amount != "" {
		v.Set(prefix+"Amount", o.Amount)
	}
	if o.CurrencyCode != "" {
		v.Set(prefix+"CurrencyCode", o.CurrencyCode)
	}

	return v
}

type Price struct {
	CurrencyCode string
	Amount       string
}

func (p Price) AddValues(prefix string, v url.Values) url.Values {
	if p.Amount != "" {
		v.Set(prefix+"Amount", p.Amount)
	}
	if p.CurrencyCode != "" {
		v.Set(prefix+"CurrencyCode", p.CurrencyCode)
	}

	return v
}

type SellerOrderAttributes struct {
	SellerOrderId     string
	StoreName         string
	CustomInformation string
}

func (s SellerOrderAttributes) AddValues(prefix string, v url.Values) url.Values {
	if s.SellerOrderId != "" {
		v.Set(prefix+"SellerOrderId", s.SellerOrderId)
	}
	if s.StoreName != "" {
		v.Set(prefix+"StoreName", s.StoreName)
	}
	if s.CustomInformation != "" {
		v.Set(prefix+"CustomInformation", s.CustomInformation)
	}

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
	Name  string
	Email string
	Phone string
}

type Destination struct {
	DestinationType     string
	PhysicalDestination Address
}

type Constraint struct {
	ConstraintID string
	Description  string
}

type OrderReferenceStatus struct {
	State               OrderReferenceState
	LastUpdateTimestamp time.Time
	ReasonCode          string
	ReasonDescription   string
}

type StatusCommon struct {
	LastUpdateTimestamp time.Time
	ReasonCode          string
	ReasonDescription   string
}

type CaptureStatus struct {
	State CaptureState
	StatusCommon
}

type AuthorizationStatus struct {
	State AuthorisationState
	StatusCommon
}

type RefundStatus struct {
	State RefundState
	StatusCommon
}

type OrderReferenceAttributes struct {
	PlatformId            string
	SellerNote            string
	OrderTotal            OrderTotal
	SellerOrderAttributes SellerOrderAttributes
}

func (o OrderReferenceAttributes) AddValues(prefix string, v url.Values) url.Values {
	if o.PlatformId != "" {
		v.Set(prefix+"PlatformId", o.PlatformId)
	}
	if o.SellerNote != "" {
		v.Set(prefix+"SellerNote", o.SellerNote)
	}

	v = o.OrderTotal.AddValues(prefix+"OrderTotal.", v)

	return o.SellerOrderAttributes.AddValues(prefix+"SellerOrderAttributes.", v)
}

type OrderReferenceDetails struct {
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
	Buyer                  Buyer
	OrderTotal             OrderTotal
	Destination            Destination
	OrderReferenceStatus   OrderReferenceStatus
	SellerOrderAttributes  SellerOrderAttributes
}

type AuthorizationDetails struct {
	AmazonAuthorizationId       string
	AuthorizationBillingAddress Address
	AuthorizationReferenceId    string
	SellerAuthorizationNote     string
	AuthorizationAmount         Price
	CapturedAmount              Price
	AuthorizationFee            Price
	IdList                      []string
	CreationTimestamp           time.Time
	ExpirationTimestamp         time.Time
	AuthorizationStatus         AuthorizationStatus
	SoftDecline                 bool
	CaptureNow                  bool
	SoftDescriptor              string
}

type BillingAgreementAttributes struct {
	PlatformId                       string
	SellerNote                       string
	SellerBillingAgreementAttributes SellerBillingAgreementAttributes
}

type SellerBillingAgreementAttributes struct {
	SellerBillingAgreementId string
	StoreName                string
	CustomInformation        string
}

type BillingAgreementDetails struct {
	AmazonBillingAgreementId         string
	BillingAddress                   Address
	SellerNote                       string
	PlatformId                       string
	ReleaseEnvironment               string
	Constraints                      []Constraint
	CreationTimestamp                time.Time
	BillingAgreementConsent          bool
	Buyer                            Buyer
	Destination                      Destination
	BillingAgreementLimits           BillingAgreementLimits
	BillingAgreementStatus           BillingAgreementStatus
	SellerBillingAgreementAttributes SellerBillingAgreementAttributes
}

type BillingAgreementLimits struct {
	AmountLimitPerTimePeriod Price
	TimePeriodStartDate      time.Time
	TimePeriodEndDate        time.Time
	CurrentRemainingBalance  Price
}

type BillingAgreementStatus struct {
	State                BillingAgreementState
	LastUpdatedTimestamp time.Time
	ReasonCode           string
	ReasonDescription    string
}

type CaptureDetails struct {
	AmazonCaptureId    string
	CaptureReferenceId string
	SellerCaptureNote  string
	CaptureAmount      Price
	RefundedAmount     Price
	CaptureFee         Price
	IdList             []string
	CreationTimestamp  time.Time
	CaptureStatus      CaptureStatus
	SoftDescriptor     string
	ConvertedAmount    Price
	ConversionRate     string
}

type RefundDetails struct {
	AmazonRefundId    string
	RefundReferenceId string
	SellerRefundNote  string
	RefundType        string
	RefundAmount      Price
	FeeRefunded       Price
	CreationTimestamp time.Time
	RefundStatus      RefundStatus
	SoftDescriptor    string
	ConvertedAmount   Price
	ConversionRate    string
}

type Message struct {
	Locale string
	Text   string
}

type GetServiceStatusResult struct {
	Status    string
	Timestamp time.Time
	MessageId string
	Messages  []Message
}
