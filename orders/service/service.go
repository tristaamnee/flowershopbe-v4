package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
	"github.com/tristaamne/flowershopbe-v4/orders/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type service struct {
	repo  repository.OrderRepository
	payOS payment.PaymentProvider
	cfg   *config.Config
}

func NewService(repo repository.OrderRepository, payOS payment.PaymentProvider, cfg *config.Config) Service {
	return &service{repo: repo, payOS: payOS, cfg: cfg}
}

type Service interface {
	Checkout(ctx context.Context, req model.OrderRequest) (map[string]interface{}, error)
}

func (s *service) Checkout(ctx context.Context, req model.OrderRequest) (map[string]interface{}, error) {

	if len(req.OrderItems) == 0 {
		return nil, fmt.Errorf("Cart is empty")
	}

	if len(req.DeliveryAddress.Phone) < 5 {
		return nil, fmt.Errorf("Invalid phone number")
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

	cancelURL := s.cfg.FEAddr + "/cancel"
	returnURL := s.cfg.FEAddr + "/return"

	rawSignature := fmt.Sprintf("amount=%d&cancelUrl=%s&description=%s&orderCode=%d&returnUrl=%s", totalPrice, cancelURL, orderDescription, orderCode, returnURL)
	checkSumKey := s.cfg.PayOSChecksum

	signature := utils.ComputeHmac256(rawSignature, checkSumKey)

	order := &model.Order{
		OrderNumber:      orderCode,
		OrderItems:       req.OrderItems,
		PromotionIDs:     req.PromotionIDs,
		TotalPrice:       totalPrice,
		DeliveryAddress:  req.DeliveryAddress,
		Status:           model.StatusPending,
		PaymentSignature: signature,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	_, err := s.repo.CreateAOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("error when create an order: %w", err)
	}

	payOSReq := map[string]interface{}{
		"orderCode":   orderCode,
		"amount":      totalPrice,
		"description": orderDescription,
		"cancelUrl":   cancelURL,
		"returnUrl":   returnURL,
		"expiredAt":   time.Now().Add(time.Minute * 30).Unix(),
		"signature":   signature,
	}

	result, err := s.payOS.CreatePaymentLink(payOSReq)
	if err != nil {
		filter := bson.M{"order_number": orderCode}
		update := bson.M{"$set": bson.M{
			"status":     model.StatusCancelled,
			"updated_at": time.Now(),
		}}
		err = s.repo.UpdateAOrder(ctx, filter, update)
		if err != nil {
			return nil, fmt.Errorf("error when update an order's status: %w", err)
		}

		return nil, fmt.Errorf("error when create a payment link: %w", err)
	}

	delete(result, "checkoutUrl")
	delete(result, "paymentLinkId")

	return result, nil

}
