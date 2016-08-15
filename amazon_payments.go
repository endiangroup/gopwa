package gopwa

import (
	"net/http"
	"net/url"
	"strings"
)

type AmazonPayments struct {
	AccessKeyID string
	SellerID    string
	HttpClient  *http.Client
	Signatory   Signatory
	Endpoint    *url.URL
	Version     string
}

func New(sellerID, accessKeyID, accessKeySecret string, region Region, environment Environment) *AmazonPayments {
	if environment == "" {
		environment = Sandbox
	}

	return &AmazonPayments{
		AccessKeyID: accessKeyID,
		SellerID:    sellerID,
		Endpoint: &url.URL{
			Scheme: "https",
			Host:   Regions[region],
			Path:   strings.Join([]string{string(environment), V20130101}, "/"),
		},
		HttpClient: http.DefaultClient,
		Signatory: V2Hmac256Signatory{
			secret: accessKeySecret,
		},
		Version: V20130101,
	}
}
