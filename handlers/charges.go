package handlers

import (
	"net/http"
	"os"
	"stripe/models"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func ChargeEUR(c *gin.Context) {
	// we will bind our JSON body to the `json` var
	var json models.ChargeJSON
	c.BindJSON(&json)
	
	// Set Stripe API key
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey
	
	// Attempt to make the charge.
	// We are setting the charge response to _
	// as we are not using it.
	_, err := charge.New(&stripe.ChargeParams{
	  Amount:       stripe.Int64(json.Amount),
	  Currency:     stripe.String(string(stripe.CurrencyEUR)),
	  Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")}, // this should come from clientside
	  ReceiptEmail: stripe.String(json.ReceiptEmail)})
	
	if err != nil {
	  // Handle any errors from attempt to charge
	  c.String(http.StatusBadRequest, "Request failed")
	  return
	}
	
	c.String(http.StatusCreated, "Successfully charged")
	}