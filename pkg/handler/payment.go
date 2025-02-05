package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func CreatePaymentIntent(c *gin.Context) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	var request struct {
		Amount int64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(request.Amount),
		Currency: stripe.String("rub"),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании платежа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_secret": intent.ClientSecret})
}
