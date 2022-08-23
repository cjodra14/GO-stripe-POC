package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/product"
)

func GetProducts(c *gin.Context) {
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey
	// products := make([]Product, 0)

	productSlice := []stripe.Product{}

	params := &stripe.ProductListParams{}
	params.Single = true
	i := product.List(params)
	for i.Next() {
		c := i.Product()
		product := stripe.Product{
			Active:              c.Active,
			Attributes:          c.Attributes,
			Caption:             c.Caption,
			Created:             c.Created,
			DeactivateOn:        c.DeactivateOn,
			Description:         c.Description,
			ID:                  c.ID,
			Images:              c.Images,
			Livemode:            c.Livemode,
			Metadata:            c.Metadata,
			Name:                c.Name,
			PackageDimensions:   c.PackageDimensions,
			Shippable:           c.Shippable,
			StatementDescriptor: c.StatementDescriptor,
			Type:                c.Type,
			UnitLabel:           c.UnitLabel,
			URL:                 c.URL,
			Updated:             c.Updated,
		}
		productSlice = append(productSlice, product)
	}
	c.JSON(http.StatusOK, productSlice)
}
