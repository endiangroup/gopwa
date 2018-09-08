# Go PayWithAmazon

[![GoDoc](https://godoc.org/github.com/endiangroup/gopwa?status.svg)](https://godoc.org/github.com/endiangroup/gopwa) [![Go Report Card](https://goreportcard.com/badge/github.com/endiangroup/gopwa)](https://goreportcard.com/report/github.com/endiangroup/gopwa)

Go PayWithAmazon (gopwa) is an api wrapper around the Amazon Payments endpoints. No additional dependencies are required to use this library.

[See the official Amazon Payments docs for more information](https://payments.amazon.co.uk/developer/documentation/lpwa/201909330)

## Features:
- Comprehensive suite of native types for all requests and responses for non-recurring endpoints (08/09/2018)
- Supports both sandbox and live endpoints
- Supports all amazon payment regions (17/08/2016)
- Error handling with native type
- New or custom request/responses are supported via interfaces (see `paywithamazon.go: PayWithAmazon.Do()`)
- Optionally supports retrying throttled requests with exponential backoff (disabled by default, enable with: `PayWithAmazon.HandleThrottling(true)`)

## Example

Construct new `PayWithAmazon` object:
```go
pwa := gopwa.New("my-seller-id", "my-access-key", "my-access-secret", gopwa.UK, gopwa.Sandbox)
```

Get an orders reference details:
```go
orderRefResp, err := pwa.GetOrderReferenceDetails(anOrderReferenceId, "")
if err != nil {
	return err
}

fmt.Printf("Request ID: %s", orderRefResp.Metadata.RequestId)

fmt.Printf("Order Reference ID: %s", orderRefResp.Result.AmazonOrderReferenceId)
fmt.Printf("Seller Note: %s", orderRefResp.Result.SellerNote)
```

Error handling (note this is very crude):
```go
func getServiceStatus(pwa PayWithAmazon) (string, error) {
	serviceStatusResp, err := pwa.GetServiceStatus()
	if err != nil {
		log.Printf("getServiceStatus: %s", err.Error())

		if errResp, ok := err.(*gopwa.ErrorResponse); ok {
			if errResp.StatusCode == 503 {
				// Backoff and retry request
				time.Sleep(1 * time.Second)
				return getServiceStatus(pwa)
			}
		}

		return "", err
	}

	return serviceStatusResp.Result.Status, nil
}
```

## Tests

### Unit Tests

Simply run `go test github.com/endiangroup/gopwa` to run all unit tests

### Integration Tests

The integration tests are a little tricker to run due to 3 factors:

1. You need an Amazon Order Reference Id to initiate an order flow, which requires using the javascript libraries they provide, logging in with OAuth and 'paying'
2. Amazon actively detects and prevents attempts to automate the frontend flow described above (no casperjs to the rescue)
3. Amazon is efficient with their allocation of Amazon Order Reference Id's. If you generate an id via the frontend and do nothing with it, re-running the frontend flow will produce the same id.

Due to these factors the integration tests must be ran one-by-one. Happy to accept any PR's that can further automate this flow.

To run the integration tests you need to do some setup:

1. You need to login to your Amazon Seller Central account and setup a sandbox user
2. You need to note down your Client ID, Seller ID, Access Key ID and Access Key Secret

To run the tests open 2 terminals, in the first:

```bash
$ go run $GOPATH/src/github.com/endiangroup/gopwa/integration/cmd/main.go -client-id={Client ID} -seller-id={Seller ID}
```

This will start a http service at `http://localhost:5000/`:

1. Browse to [http://localhost:5000/](http://localhost:5000/)
2. Click the `Pay With Amazon` button
3. Login with your sandbox user
4. Wait until both widgets have loaded ~1-2seconds
5. Click the `Pay Now` button
6. Check the output of your terminal running the http service, you should see a Amazon Order Reference Id like `S03-1673721-7848334` in your Stdout

In the second terminal run:

```bash
$ go test -tags integration $GOPATH/src/github.com/endiangroup/gopwa/ -seller-id={Seller ID} -key-id={Access Key ID} -key-secret={Access Key Secret} -{test-name}={Amazon Order Reference ID}
```

Where `{test-name}` is replace with one of the integration test flag names, you can see them with the following command:

```bash
$ cat $GOPATH/src/github.com/endiangroup/gopwa/integration_test.go | grep 'flag\.String'
```

Examples:

```bash
$ go test -tags integration $GOPATH/src/github.com/endiangroup/gopwa/ -seller-id={Seller ID} -key-id={Access Key ID} -key-secret={Access Key Secret} -authorize={Amazon Order Reference ID}
$ go test -tags integration $GOPATH/src/github.com/endiangroup/gopwa/ -seller-id={Seller ID} -key-id={Access Key ID} -key-secret={Access Key Secret} -capture={Amazon Order Reference ID}
$ go test -tags integration $GOPATH/src/github.com/endiangroup/gopwa/ -seller-id={Seller ID} -key-id={Access Key ID} -key-secret={Access Key Secret} -close-order-reference={Amazon Order Reference ID}
```

## TODO:
- Validation of request parameters (add a `Validate()` method to `Request` interface?)
- Automate integration tests
- ~~Optional retry handling of failed requests (see official python library for example)~~
- Add recurring payment specific helpers, requests and responses
