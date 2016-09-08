package integration

type Credentials struct {
	ClientID string
	SellerID string
}

type ButtonModel struct {
	ButtonType     string
	ButtonColour   string
	ButtonSize     string
	ButtonLanguage string
	Scopes         string
	PopupParam     string
	RedirectURL    string
	Credentials
}

func NewButtonModel(widgetURL, clientID, sellerID string) *ButtonModel {
	return &ButtonModel{
		Credentials: Credentials{
			ClientID: clientID,
			SellerID: sellerID,
		},
		ButtonType:     "PwA",
		ButtonColour:   "Gold",
		ButtonSize:     "medium",
		ButtonLanguage: "en-GB",
		Scopes:         "profile payments:widget payments:shipping_address payments:billing_address",
		PopupParam:     "true",
		RedirectURL:    widgetURL,
	}
}

type WidgetModel struct {
	Credentials
}

func NewWidgetModel(clientID, sellerID string) *WidgetModel {
	return &WidgetModel{
		Credentials: Credentials{
			ClientID: clientID,
			SellerID: sellerID,
		},
	}
}
