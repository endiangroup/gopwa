// +build !integration

package gopwa

import (
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ItDefaultsToSandboxIfEmptyEnvironmentPassed(t *testing.T) {
	pwa := New("", "", "", UK, "")

	expectedPath := Sandbox + "/2013-01-01"
	assert.Equal(t, string(expectedPath), pwa.Endpoint.Path)
}

func Test_ItStripsContentLengthHeader(t *testing.T) {
	mockRequest := setupMockRequest(url.Values{})

	var wasCalled bool
	assertEncodedForm := func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true
		w.WriteHeader(500)

		assert.Equal(t, "", r.Header.Get("Content-Length"))
	}

	client, server := setupTestHttp(assertEncodedForm)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)

	pwa.Do(mockRequest, &map[string]interface{}{})

	assert.True(t, wasCalled)
}

func Test_ItAddsContentTypeHeader(t *testing.T) {
	mockRequest := setupMockRequest(url.Values{})

	var wasCalled bool
	assertEncodedForm := func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true
		w.WriteHeader(500)

		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
	}

	client, server := setupTestHttp(assertEncodedForm)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)

	pwa.Do(mockRequest, map[string]interface{}{})

	assert.True(t, wasCalled)
}

func Test_ItSetsAppropriateUserAgentHeader(t *testing.T) {
	mockRequest := setupMockRequest(url.Values{})

	var wasCalled bool
	assertEncodedForm := func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true
		w.WriteHeader(500)

		assert.Equal(t, UserAgent+"/"+Version, r.Header.Get("User-Agent"))
	}

	client, server := setupTestHttp(assertEncodedForm)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)

	pwa.Do(mockRequest, map[string]interface{}{})

	assert.True(t, wasCalled)
}

func Test_ItEncodesAmazonRequestIntoUrlEncodedForm(t *testing.T) {
	action := "MockRequest"
	awsAccessKeyId := "bcd234"
	sellerID := "abc123"

	mockRequest := setupMockRequest(url.Values{
		"T1": []string{"a"},
		"T2": []string{"b"},
	})

	sigMethod := "Mock256"
	sigVersion := "1"
	sig := "abcdefghijkl"
	mockSignatory := &MockSignatory{}
	mockSignatory.On("Method").Return(sigMethod)
	mockSignatory.On("Version").Return(sigVersion)
	mockSignatory.On("Sign", mock.AnythingOfType("string")).Return([]byte(sig))

	NowForce(time.Now())
	defer NowReset()

	expectedForm := url.Values{
		"Action":           []string{action},
		"AWSAccessKeyId":   []string{awsAccessKeyId},
		"SellerId":         []string{sellerID},
		"Timestamp":        []string{Now().UTC().Format("2006-01-02T15:04:05Z")},
		"Version":          []string{string(V20130101)},
		"SignatureMethod":  []string{sigMethod},
		"SignatureVersion": []string{sigVersion},
		"Signature":        []string{base64.StdEncoding.EncodeToString([]byte(sig))},
		"T1":               []string{"a"},
		"T2":               []string{"b"},
	}

	var wasCalled bool
	assertEncodedForm := func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true
		w.WriteHeader(500)

		assert.NoError(t, r.ParseForm())

		assert.Equal(t, expectedForm, r.Form)
	}

	client, server := setupTestHttp(assertEncodedForm)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)
	pwa.Signatory = mockSignatory

	pwa.Do(mockRequest, map[string]interface{}{})

	assert.True(t, wasCalled)
}

func Test_ItDecodesResponseBodyIntoResponseObject(t *testing.T) {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}

	mockRequest := setupMockRequest(url.Values{})
	expectedResponse := &Result{
		XMLName: xml.Name{
			Space: "",
			Local: "Person",
		},
		Name:  "Grace R. Emlin",
		Phone: "",
		Email: []Email{
			Email{
				Where: "home",
				Addr:  "gre@example.com",
			},
			{
				Where: "work",
				Addr:  "gre@work.com",
			},
		},
		Groups: []string{
			"Friends", "Squash",
		},
		Address: Address{
			City:  "Hanga Roa",
			State: "Easter Island",
		},
	}

	var wasCalled bool
	assertResponse := func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true

		w.Write([]byte(`
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
			<Value>Friends</Value>
			<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
		`))
	}

	client, server := setupTestHttp(assertResponse)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)

	resp := &Result{}
	assert.NoError(t, pwa.Do(mockRequest, resp))

	assert.True(t, wasCalled)
	assert.Equal(t, expectedResponse, resp)
}

func Test_ItReturnsDecodedResponseErrorOnRequestError(t *testing.T) {
	mockRequest := setupMockRequest(url.Values{})
	expectedError := &ErrorResponse{
		StatusCode: 400,
		RequestID:  "b7afc6c3-6f75-4707-bcf4-0475ad23162c",
		Errors: []Error{
			Error{
				Type:    "Sender",
				Code:    "InvalidClientTokenId",
				Message: " The AWS Access Key Id you provided does not exist in our records. ",
				Detail:  "com.amazonservices.mws.model.Error$Detail@17b6643",
			},
			Error{
				Type:    "Sender",
				Code:    "InvalidSignature",
				Message: " The signature you provided does not match the expected. ",
				Detail:  "com.amazonservices.mws.model.Error$Detail@17b6643",
			},
		},
	}

	var wasCalled bool
	responseError := func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true

		w.WriteHeader(400)
		w.Write([]byte(`
		<ErrorResponse xmlns="http://mws.amazonservices.com/doc/2009-01-01/">
			<Error>
				<Type>Sender</Type>
				<Code>InvalidClientTokenId</Code>
				<Message> The AWS Access Key Id you provided does not exist in our records. </Message>
				<Detail>com.amazonservices.mws.model.Error$Detail@17b6643</Detail>
			</Error>
			<Error>
				<Type>Sender</Type>
				<Code>InvalidSignature</Code>
				<Message> The signature you provided does not match the expected. </Message>
				<Detail>com.amazonservices.mws.model.Error$Detail@17b6643</Detail>
			</Error>
			<RequestID>b7afc6c3-6f75-4707-bcf4-0475ad23162c</RequestID>
		</ErrorResponse>
		`))
	}

	client, server := setupTestHttp(responseError)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)

	err := pwa.Do(mockRequest, map[string]interface{}{})

	assert.True(t, wasCalled)
	assert.Equal(t, expectedError, err)
}

func Test_ItRetriesFaild5XXRequestsWithBackoffDurationWhenHandleThrottlingIsTrue(t *testing.T) {
	mockRequest := setupMockRequest(url.Values{})
	BackoffDurations = []time.Duration{
		0 * time.Millisecond,
		25 * time.Millisecond,
		50 * time.Millisecond,
		75 * time.Millisecond,
	}

	callTimestamps := []time.Time{}
	var now time.Time
	responseError := func(w http.ResponseWriter, r *http.Request) {
		if len(callTimestamps) == 0 {
			now = Now()
		}
		callTimestamps = append(callTimestamps, Now())

		if len(callTimestamps)%2 == 0 {
			w.WriteHeader(503)
		} else {
			w.WriteHeader(500)
		}

		w.Write([]byte(`
		<ErrorResponse xmlns="http://mws.amazonservices.com/doc/2009-01-01/">
			<Error>
				<Type>Sender</Type>
				<Code>RequestThrottled</Code>
				<Message>The frequency of requests was greater than allowed.</Message>
				<Detail>com.amazonservices.mws.model.Error$Detail@17b6643</Detail>
			</Error>
			<RequestID>b7afc6c3-6f75-4707-bcf4-0475ad23162c</RequestID>
		</ErrorResponse>
		`))
	}

	client, server := setupTestHttp(responseError)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)
	pwa.HandleThrottling(true)

	err := pwa.Do(mockRequest, map[string]interface{}{})
	assert.NotNil(t, err)

	assert.Len(t, callTimestamps, len(BackoffDurations))

	assert.WithinDuration(t, now, callTimestamps[0], 10*time.Millisecond)
	assert.WithinDuration(t, now.Add(BackoffDurations[1]), callTimestamps[1], 10*time.Millisecond)
	assert.WithinDuration(t, now.Add(BackoffDurations[1]+BackoffDurations[2]), callTimestamps[2], 10*time.Millisecond)
	assert.WithinDuration(t, now.Add(BackoffDurations[1]+BackoffDurations[2]+BackoffDurations[3]), callTimestamps[3], 10*time.Millisecond)
}

func Test_ItDoesntRetryFaildRequestsWhenHandleThrottlingIsFalse(t *testing.T) {
	mockRequest := setupMockRequest(url.Values{})

	var numberOfCalls uint
	responseError := func(w http.ResponseWriter, r *http.Request) {
		numberOfCalls++

		w.WriteHeader(503)
		w.Write([]byte(`
		<ErrorResponse xmlns="http://mws.amazonservices.com/doc/2009-01-01/">
		</ErrorResponse>
		`))
	}

	client, server := setupTestHttp(responseError)
	defer server.Close()

	pwa := setupAmazonPayments(client, server.URL)
	pwa.HandleThrottling(false)

	err := pwa.Do(mockRequest, map[string]interface{}{})
	assert.NotNil(t, err)

	assert.True(t, numberOfCalls == 1)
}

func setupTestHttp(serverFn func(w http.ResponseWriter, r *http.Request)) (*http.Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(serverFn))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	return &http.Client{Transport: transport}, server
}

func setupMockRequest(v url.Values) *MockRequest {
	mockRequest := &MockRequest{}
	mockRequest.On("Action").Return("MockRequest")
	mockRequest.On("AddValues", url.Values{}).Return(v)

	return mockRequest
}

func setupAmazonPayments(client *http.Client, urlString string) *PayWithAmazon {
	pwa := New("abc123", "bcd234", "", UK, "")
	pwa.Endpoint, _ = url.Parse(urlString)
	pwa.HttpClient = client

	return pwa
}
