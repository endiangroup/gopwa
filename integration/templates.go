package integration

import "html/template"

var ButtonTemplate = template.Must(template.New("button").Parse(`
<body>
// Add to the body of the page:
<div id="LoginWithAmazon"> </div>  
<script>
window.onAmazonLoginReady = function() { 
	amazon.Login.setClientId('{{.ClientID}}'); 
};

window.onAmazonPaymentsReady = function(){
	// render the button here
	var authRequest; 

	OffAmazonPayments.Button("LoginWithAmazon", "{{.SellerID}}", { 
		type:  "{{.ButtonType}}", 
		color: "{{.ButtonColour}}", 
		size:  "{{.ButtonSize}}", 
		language: "{{.ButtonLanguage}}",

		authorization: function() { 
			loginOptions = {scope: "{{.Scopes}}", popup: "{{.PopupParam}}"}; 
			authRequest = amazon.Login.authorize (loginOptions, "{{.RedirectURL}}"); 
		}, 

		onError: function(error) { 
			// your error handling code.
			alert("The following error occurred: " 
			       + error.getErrorCode() 
			       + ' - ' + error.getErrorMessage());
		} 
	});
}
</script>

<script async="async" type='text/javascript' src='https://static-eu.payments-amazon.com/OffAmazonPayments/gbp/sandbox/lpa/js/Widgets.js'>
</script>
<script type="text/javascript">
document.getElementById('Logout').onclick = function() {
	amazon.Login.logout();
};
</script>
</body>
`))

var WidgetTemplate = template.Must(template.New("widget").Parse(`
<body>
	<button type="button" id="payNow" style="
    width: 50%;
    height: 50px;
    margin: 20px 25%;
">Pay Now</button>

	<div id="addressBookWidgetDiv"> </div>

	<div id="walletWidgetDiv"> </div>

	<script>
	var requestData = new Array();

	window.onAmazonLoginReady = function() {amazon.Login.setClientId('{{.ClientID}}'); };
	window.onAmazonPaymentsReady = function() {
		new OffAmazonPayments.Widgets.AddressBook({
			sellerId: '{{.SellerID}}',
			onOrderReferenceCreate: function(orderReference) {
				requestData.push('orderReferenceID='+orderReference.getAmazonOrderReferenceId());
			},
			onAddressSelect: function(orderReference) {
				// Replace the following code with the action that you want
				// to perform after the address is selected. The 
				// amazonOrderReferenceId can be used to retrieve the address 
				// details by calling the GetOrderReferenceDetails operation. 

				// If rendering the AddressBook and Wallet widgets
				// on the same page, you do not have to provide any additional
				// logic to load the Wallet widget after the AddressBook widget.
				// The Wallet widget will re-render itself on all subsequent 
				// onAddressSelect events, without any action from you. 
				// It is not recommended that you explicitly refresh it.
			},
			design: {
				designMode: 'responsive'
			},
			onReady: function(orderReference) {
				// Enter code here you want to be executed 
				// when the address widget has been rendered. 
			},
			onError: function(error) {
				requestData.push('addressError='+encodeURIComponent(error.getErrorCode() + ': ' + error.getErrorMessage()));
			}
		}).bind("addressBookWidgetDiv");

		new OffAmazonPayments.Widgets.Wallet({
			sellerId: '{{.SellerID}}',
			onPaymentSelect: function(orderReference) {
				// Replace this code with the action that you want to perform
				// after the payment method is selected.

				// Ideally this would enable the next action for the buyer
				// including either a "Continue" or "Place Order" button.
			},
			design: {
				designMode: 'responsive'
			},
			onError: function(error) {
				requestData.push('walletError='+encodeURIComponent(error.getErrorCode() + ': ' + error.getErrorMessage()));
			}
		}).bind("walletWidgetDiv");
	};

	document.getElementById("payNow").onclick = function () { 
		var request = new XMLHttpRequest();
		request.open("POST", "/callback", true);
		request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
		request.send(requestData.join('&').replace(/%20/, '+'));

		window.location = window.location.protocol + "//" + window.location.host;
	};
	</script>

	<script async="async" 
		src='https://static-eu.payments-amazon.com/OffAmazonPayments/uk/sandbox/lpa/js/Widgets.js'>
	</script>
</body>
`))
