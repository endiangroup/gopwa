# Go PayWithAmazon

Go PayWithAmazon (gopwa) is an api wrapper around the Amazon Payments endpoints. No additional dependencies are required to use this library.

[See the official Amazon Payments docs for more information](https://payments.amazon.co.uk/developer/documentation/lpwa/201909330)

## Features:
- Comprehensive suite of native types for all requests and responses (17/08/2016)
- Supports both sandbox and live endpoints
- Supports all known amazon payment regions (17/08/2016)
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

## TODO:
- Validation of request parameters (add a `Validate()` method to `Request` interface?)
- ~~Optional retry handling of failed requests (see official python library for example)~~
