package main

import (
	"log"
	"stripe/handlers"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.Default())

	//Charges
	r.POST("/api/charges", handlers.ChargeEUR)
	r.POST("api/charges/intent", handlers.PaymentIntent)

	//Customer 
	r.POST("/api/customer", handlers.CreateCustomer)
	r.PUT("/api/customer", handlers.UpdateCustomer)
	r.GET("api/customer", handlers.GetCustomers)
	r.GET("api/customer/id", handlers.GetCustomersByID)
	r.GET("api/customer/email", handlers.GetCustomersByEmail)
	r.DELETE("api/customer", handlers.DeleteCustomerByID)

	//Product
	r.GET("api/product", handlers.GetProducts)

	//Subscription 
	r.GET("api/subscription", handlers.GetSubscriptions)

	r.Run(":8413")
}
