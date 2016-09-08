package integration

import (
	"errors"
	"html/template"
	"net/http"
)

type IntegrationHandler struct {
	ButtonTemplate *template.Template
	ButtonModel    *ButtonModel

	WidgetTemplate *template.Template
	WidgetModel    *WidgetModel

	ErrorsCh   chan []error
	OrderRefCh chan string
}

func New(widgetURL, clientID, sellerID string, buttonTemplate, widgetTemplate *template.Template) *IntegrationHandler {
	return &IntegrationHandler{
		ButtonTemplate: buttonTemplate,
		WidgetTemplate: widgetTemplate,
		ButtonModel:    NewButtonModel(widgetURL, clientID, sellerID),
		WidgetModel:    NewWidgetModel(clientID, sellerID),
		ErrorsCh:       make(chan []error, 1),
		OrderRefCh:     make(chan string, 1),
	}
}

func (app *IntegrationHandler) ButtonHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.ButtonTemplate.Execute(w, app.ButtonModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *IntegrationHandler) WidgetHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.WidgetTemplate.Execute(w, app.WidgetModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *IntegrationHandler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	if r.PostForm.Get("addressError") != "" || r.PostForm.Get("walletError") != "" {
		errs := []error{}
		if r.PostForm.Get("addressError") != "" {
			errs = append(errs, errors.New(r.PostForm.Get("addressError")))
		}
		if r.PostForm.Get("walletError") != "" {
			errs = append(errs, errors.New(r.PostForm.Get("walletError")))
		}

		app.ErrorsCh <- errs
	}

	if r.PostForm.Get("orderReferenceID") != "" {
		app.OrderRefCh <- r.PostForm.Get("orderReferenceID")
	}
}
