package main

import (
	"log"
	"stripe/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// ChargeJSON incoming data for Stripe API

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set up server
	r := gin.Default()

	// our basic charge API route
	r.POST("/api/charges", handlers.ChargeEUR)

	r.POST("/api/customer", handlers.CreateCustomer)
	r.PUT("/api/customer", handlers.UpdateCustomer)
	r.GET("api/customer", handlers.GetCustomers)
	r.GET("api/customer/email", handlers.GetCustomersByEmail)
	r.DELETE("api/customer",handlers.DeleteCustomerByID)

	r.Run(":8413")
}
