package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
)

func GetSubscriptions(c *gin.Context) {
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	subscriptionSlice := []stripe.Subscription{}
	params := &stripe.SubscriptionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := sub.List(params)
	for i.Next() {
		s := i.Subscription()
		subscriptionData := stripe.Subscription{
			ApplicationFeePercent:         s.ApplicationFeePercent,
			BillingCycleAnchor:            s.BillingCycleAnchor,
			BillingThresholds:             s.BillingThresholds,
			CancelAt:                      s.CancelAt,
			CancelAtPeriodEnd:             s.CancelAtPeriodEnd,
			CanceledAt:                    s.CanceledAt,
			CollectionMethod:              s.CollectionMethod,
			Created:                       s.Created,
			CurrentPeriodEnd:              s.CurrentPeriodEnd,
			CurrentPeriodStart:            s.CurrentPeriodStart,
			Customer:                      s.Customer,
			DaysUntilDue:                  s.DaysUntilDue,
			DefaultPaymentMethod:          s.DefaultPaymentMethod,
			DefaultSource:                 s.DefaultSource,
			DefaultTaxRates:               s.DefaultTaxRates,
			Discount:                      s.Discount,
			EndedAt:                       s.EndedAt,
			ID:                            s.ID,
			Items:                         s.Items,
			LatestInvoice:                 s.LatestInvoice,
			Livemode:                      s.Livemode,
			Metadata:                      s.Metadata,
			NextPendingInvoiceItemInvoice: s.NextPendingInvoiceItemInvoice,
			Object:                        s.Object,
			OnBehalfOf:                    s.OnBehalfOf,
			PauseCollection:               s.PauseCollection,
			PendingInvoiceItemInterval:    s.PendingInvoiceItemInterval,
			PendingSetupIntent:            s.PendingSetupIntent,
			PendingUpdate:                 s.PendingUpdate,
			Plan:                          s.Plan,
			Quantity:                      s.Quantity,
			Schedule:                      s.Schedule,
			StartDate:                     s.StartDate,
			Status:                        s.Status,
			TransferData:                  s.TransferData,
			TrialEnd:                      s.TrialEnd,
			TrialStart:                    s.TrialStart,
			TaxPercent:                    s.TaxPercent,
		}
		subscriptionSlice = append(subscriptionSlice, subscriptionData)
	}
	c.JSON(http.StatusOK, subscriptionSlice)
}
