package handlers

import (
	"fmt"
	"net/http"
	"os"
	"stripe/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func CreateCustomer(c *gin.Context) {
	var json models.CustomerJSON
	c.BindJSON(&json)

	// Set Stripe API key
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	fmt.Println(json)

	params := &stripe.CustomerParams{
		Email: &json.Email,
		Name:  &json.Name,
	}

	createdCustomer, err := customer.New(params)
	if err != nil {
		c.String(http.StatusBadRequest, "Request failed")
		return
	}
	fmt.Println(createdCustomer)
	c.JSON(http.StatusOK, createdCustomer)
}

func GetCustomers(c *gin.Context) {

	type Customers struct {
		Customers []stripe.Customer `json:"customers,omitempty"`
	}
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	customers := Customers{}
	customerSlice := []stripe.Customer{}

	params := &stripe.CustomerListParams{}
	params.Single = true
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		customerInfo := stripe.Customer{
			ID:       c.ID,
			Name:     c.Name,
			Email:    c.Email,
			Address:  c.Address,
			Balance:  c.Balance,
			Created:  c.Created,
			Currency: c.Currency,
		}
		customerSlice = append(customerSlice, customerInfo)
	}
	// customers.Customers = customerSlice
	fmt.Println(customers)

	// customerBytes, err := json.Marshal(customers)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Request failed")
	// 	return
	// }
	c.JSON(http.StatusOK, customerSlice)
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

func GetCustomersByEmail(c *gin.Context) {
	var customerSelectionJSON models.CustomerJSON
	c.BindJSON(&customerSelectionJSON)

	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	params := &stripe.CustomerListParams{
		Email: &customerSelectionJSON.Email,
	}

	params.Single = true

	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		fmt.Println(c.ID)
		fmt.Println(c.Email)
	}

	c.JSON(http.StatusOK, i.Customer())
}

func GetCustomersByID(c *gin.Context) {
	var customerSelectionJSON models.CustomerJSON
	c.BindJSON(&customerSelectionJSON)

	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	customerData, err := customer.Get(customerSelectionJSON.ID, nil)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(http.StatusOK, customerData)
}

func DeleteCustomerByID(c *gin.Context) {
	var json models.CustomerJSON
	c.BindJSON(&json)

	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	customer, err := customer.Del(json.ID, nil)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(http.StatusOK, customer)
}
