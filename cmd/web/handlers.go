package main

import (
	"net/http"
)

func (app *application) Payment(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSuccess(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.errorLog.Println(err)
	}
	name := r.Form.Get("cardholder_name")
	email := r.Form.Get("cardholder_email")
	payment_intent := r.Form.Get("payment_intent")
	payment_method := r.Form.Get("payment_method")
	payment_amount := r.Form.Get("payment_amount")
	payment_currency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["name"] = name
	data["email"] = email
	data["payment_amount"] = payment_amount
	data["payment_intent"] = payment_intent
	data["payment_method"] = payment_method
	data["payment_currency"] = payment_currency

	if err := app.renderTemplate(w, r, "success", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
