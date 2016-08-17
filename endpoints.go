package gopwa

type Region uint8

const (
	UK Region = iota
	DE
	EU
	US
	NA
	JP
)

var Regions = []string{
	UK: "mws-eu.amazonservices.com",
	DE: "mws-eu.amazonservices.com",
	EU: "mws-eu.amazonservices.com",
	US: "mws.amazonservices.com",
	NA: "mws.amazonservices.com",
	JP: "mws.amazonservices.jp",
}

type ApiVersion string

const (
	V20130101 ApiVersion = "2013-01-01"
)

type Environment string

const (
	Sandbox Environment = "/OffAmazonPayments_Sandbox"
	Live                = "/OffAmazonPayments"
)
