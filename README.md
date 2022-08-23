First step  create .env file with the next field:

SK_TEST_KEY=( available on the stripe dashboard)

To run the back -->

go run main.go


Possible cURL's


-> Customers

Create new Customer:

```
curl --header "Content-Type: application/json" \
  --request POST \                 
  --data '{"name":"Chris","email":"go_poc@example.com","tax_exempt":"none","balance":100000,"currency":"eur","payment_method":"pm_card_visa","invoice_settings":{"default_payment_method":"pm_card_visa"},"preferred_locales":["es","en"]}' \
  http://localhost:8413/api/customer

```

Update Customer info:

```
curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"id":"cus_MBQfADqzheFwVj","name":"Christian","email":"christian_poc_go@example.com"}' \
  http://localhost:8413/api/customer  
```

Delete Customer by ID:

```
curl --header "Content-Type: application/json" \
  --request DELETE \
  --data '{"id":"cus_MBQg9OWdDmhQwN"}' \                                                          
  http://localhost:8413/api/customer    
```


Get Customer info by Email 

```
curl --header "Content-Type: application/json" \
  --request GET \
  --data '{"email":"christian_poc_go@example.com"}' \

```

Get Customer by ID

```
curl --header "Content-Type: application/json" \
  --request GET \
  --data '{"cus_MISTvcjnQPpQlH"}' \
  http://localhost:8413/api/customer
```


Get All Customers :
```
curl --request GET \                            
 http://localhost:8413/api/customer
```

-> Payments

Make a charge:

```
 curl --header "Content-Type: application/json" \
  --request POST \                 
  --data '{"amount":5000,"receiptEmail":"go_poc@example.com"}' \                                                                  
  http://localhost:8413/api/charges 
```

Get Products 

```
curl --request GET \                            
 http://localhost:8413/api/products
```
Get Subscriptions

```
curl --request GET \                            
 http://localhost:8413/api/subscriptions
```
