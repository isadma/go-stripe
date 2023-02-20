package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) GetPaymentIntent(response http.ResponseWriter, request *http.Request) {
	j := jsonResponse{
		OK: true,
	}

	out, err := json.MarshalIndent(j, "", " ")
	if err != nil {
		app.errorLog.Println(err)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(out)
}
