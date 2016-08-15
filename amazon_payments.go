package gopwa

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	UserAgent = "Go_PayWithAmazon"
	Version   = "0.0.1"
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

func (ap AmazonPayments) buildForm(v url.Values, action, method string) url.Values {
	v.Set("Action", action)
	v.Set("AWSAccessKeyId", ap.AccessKeyID)
	v.Set("SellerId", ap.SellerID)
	v.Set("Timestamp", Now().UTC().Format("2006-01-02T15:04:05Z"))
	v.Set("Version", ap.Version)
	v.Set("SignatureMethod", ap.Signatory.Method())
	v.Set("SignatureVersion", ap.Signatory.Version())
	v.Set("Signature", base64.StdEncoding.EncodeToString(
		ap.Signatory.Sign(ap.prepareSignature(method, v)),
	))

	return v
}

func (ap AmazonPayments) Do(amazonReq Request, response interface{}) error {
	method := "POST"
	form := ap.buildForm(amazonReq.AddValues(url.Values{}), amazonReq.Action(), method)

	req, err := http.NewRequest(method, ap.Endpoint.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.ContentLength = 0
	req.Header.Del("Content-Length")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", UserAgent+"/"+Version)

	resp, err := ap.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("!")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return xml.Unmarshal(body, response)
}
