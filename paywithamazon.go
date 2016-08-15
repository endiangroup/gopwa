package gopwa

import (
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	UserAgent = "Go_PayWithAmazon"
	Version   = "0.0.1"
)

type PayWithAmazon struct {
	AccessKeyID string
	SellerID    string
	HttpClient  *http.Client
	Signatory   Signatory
	Endpoint    *url.URL
	Version     string
}

func New(sellerID, accessKeyID, accessKeySecret string, region Region, environment Environment) *PayWithAmazon {
	if environment == "" {
		environment = Sandbox
	}

	return &PayWithAmazon{
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

func (pwa PayWithAmazon) buildForm(v url.Values, action, method string) url.Values {
	v.Set("Action", action)
	v.Set("AWSAccessKeyId", pwa.AccessKeyID)
	v.Set("SellerId", pwa.SellerID)
	v.Set("Timestamp", Now().UTC().Format("2006-01-02T15:04:05Z"))
	v.Set("Version", pwa.Version)
	v.Set("SignatureMethod", pwa.Signatory.Method())
	v.Set("SignatureVersion", pwa.Signatory.Version())
	v.Set("Signature", base64.StdEncoding.EncodeToString(
		pwa.Signatory.Sign(pwa.prepareSignature(method, v)),
	))

	return v
}

func (pwa PayWithAmazon) Do(amazonReq Request, response interface{}) error {
	method := "POST"
	form := pwa.buildForm(amazonReq.AddValues(url.Values{}), amazonReq.Action(), method)

	req, err := http.NewRequest(method, pwa.Endpoint.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	resp, err := pwa.HttpClient.Do(pwa.setHeaders(req))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		responseError := &ErrorResponse{StatusCode: resp.StatusCode}

		if err := xml.Unmarshal(body, responseError); err != nil {
			return err
		}

		return responseError
	}

	return xml.Unmarshal(body, response)
}

func (pwa PayWithAmazon) setHeaders(req *http.Request) *http.Request {
	req.ContentLength = 0
	req.Header.Del("Content-Length")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", UserAgent+"/"+Version)

	return req
}
