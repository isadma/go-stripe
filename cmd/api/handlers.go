package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/isadma/go-stripe/internal/cards"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) GetPaymentIntent(response http.ResponseWriter, request *http.Request) {

	var payload stripePayload

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.secret,
		Currency: payload.Currency,
	}

	okay := true
	pi, message, err := card.Charge(payload.Currency, amount)
	if err != nil {
		okay = false
	}

	if okay {
		out, err := json.MarshalIndent(pi, "", " ")
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(out)
	} else {
		j := jsonResponse{
			OK:      false,
			Message: message,
			Content: "",
		}

		out, err := json.MarshalIndent(j, "", " ")
		if err != nil {
			app.errorLog.Println(err)
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(out)
	}
}
