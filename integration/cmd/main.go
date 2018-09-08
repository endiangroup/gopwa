package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/endiangroup/gopwa/integration"
)

var clientIDFlag = flag.String("client-id", "", "MWS Client ID")
var sellerIDFlag = flag.String("seller-id", "", "MWS Seller ID")
var portFlag = flag.String("port", "5000", "HTTP Server Port")

func main() {
	flag.Parse()

	handler := integration.New(fmt.Sprintf("http://localhost:%s/widgets", *portFlag), *clientIDFlag, *sellerIDFlag, integration.ButtonTemplate, integration.WidgetTemplate)

	http.HandleFunc("/", handler.ButtonHandler)
	http.HandleFunc("/widgets", handler.WidgetHandler)
	http.HandleFunc("/callback", handler.CallbackHandler)

	go func() {
		for {
			select {
			case orderRefID := <-handler.OrderRefCh:
				fmt.Println(orderRefID)
			case errs := <-handler.ErrorsCh:
				fmt.Println("Errors:", errs)
			}
		}
	}()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *portFlag), nil))
}
