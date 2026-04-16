package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
)

func CheckOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.OrderRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Order body": err.Error()})
			return
		}
		if len(req.OrderItems) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Cart was empty",
			})
			return
		}

		if len(req.DeliveryAddress.Phone) < 5 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid phone number",
			})
			return
		}

		var totalPrice int64
		for _, orderItem := range req.OrderItems {
			totalPrice += orderItem.Price
		}

		phoneNumLastFourNum := req.DeliveryAddress.Phone[len(req.DeliveryAddress.Phone)-4:]
		suffix, _ := strconv.ParseInt(phoneNumLastFourNum, 10, 64)
		timestamp := time.Now().Unix()
		orderCode := timestamp*10000 + suffix

		orderDescription := "Thanh-toan-don-hang-" + strconv.FormatInt(orderCode, 10)

		cancelURL := os.Getenv("FE_ADDR") + "/cancel"
		returnURL := os.Getenv("FE_ADDR") + "/return"

		rawSignature := fmt.Sprintf("amount=%d&cancelUrl=%s&description=%s&orderCode=%d&returnUrl=%s", totalPrice, cancelURL, orderDescription, orderCode, returnURL)
		checkSumKey := os.Getenv("PAYOS_CHECKSUM")

		signature := utils.ComputeHmac256(rawSignature, checkSumKey)

		payOSReq := map[string]interface{}{
			"orderCode":   orderCode,
			"amount":      totalPrice,
			"description": orderDescription,
			"cancelUrl":   cancelURL,
			"returnUrl":   returnURL,
			"expiredAt":   time.Now().Add(time.Minute * 30).Unix(),
			"signature":   signature,
		}

		payOS := payment.NewPayOSProvider()
		result, err := payOS.CreatePaymentLink(payOSReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error when create a payment link": err.Error(),
			})
			return
		}

		delete(result, "checkoutUrl")
		delete(result, "paymentLinkId")

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"payment": result["data"],
		})

	}
}
