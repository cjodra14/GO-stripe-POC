package handlers

import (
	"fmt"
	"net/http"
	"os"
	"stripe/models"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func CreateCustomer(c *gin.Context) {
	var json models.CustomerJSON
	c.BindJSON(&json)

	// Set Stripe API key
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	params := &stripe.CustomerParams{
		Email:         &json.Email,
		Name:          &json.Name,
		TaxExempt:     &json.TaxExempt,
		PaymentMethod: &json.PaymentMethod,
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: &json.InvoiceSettings.DefaultPaymentMethod,
		},
		PreferredLocales: stripe.StringSlice(json.PreferredLocales),
	}

	createdCustomer, err := customer.New(params)
	if err != nil {
		c.String(http.StatusBadRequest, "Request failed")
		return
	}
	fmt.Println(createdCustomer)
}

func GetCustomers(c *gin.Context) {
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey
	
	params := &stripe.CustomerListParams{}
	params.Single = true
	i:= customer.List(params)
	for i.Next(){
		c := i.Customer()
		fmt.Println(c.ID)
	}
}

func UpdateCustomer(c *gin.Context) {
	var json models.CustomerJSON
	c.BindJSON(&json)

	// Set Stripe API key
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	params := &stripe.CustomerParams{
		Email: &json.Email,
		Name:  &json.Name,
	}

	updatedCustomer, err := customer.Update(json.ID, params)
	if err != nil {
		c.String(http.StatusBadRequest, "Request failed")
		return
	}
	fmt.Println(updatedCustomer)
}
