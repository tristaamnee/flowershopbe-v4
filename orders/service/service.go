package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/payment"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
	"github.com/tristaamne/flowershopbe-v4/orders/repository"
	prodModel "github.com/tristaamne/flowershopbe-v4/products/model"
	prodService "github.com/tristaamne/flowershopbe-v4/products/service"
	userModel "github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	prodSrv prodService.Service
	repo    repository.OrderRepository
	payOS   payment.PaymentProvider
	cfg     *config.Config
	rdb     *redis.Client
}

func NewService(repo repository.OrderRepository, prodSrv prodService.Service, payOS payment.PaymentProvider, cfg *config.Config, rdb *redis.Client) Service {
	return &service{repo: repo, prodSrv: prodSrv, payOS: payOS, cfg: cfg, rdb: rdb}
}

type Service interface {
	MemberCheckout(ctx context.Context, req model.MemberOrderRequest, userIdStr string) (map[string]interface{}, error)
	GuestCheckout(ctx context.Context, req model.GuestOrderRequest) (map[string]interface{}, error)
	UpdateOrderStatus(ctx context.Context, orderCode int, status string) error
}

func (s *service) updateQuantityForMultiple(ctx context.Context, rollbackData map[primitive.ObjectID]interface{}, items []prodModel.Product, changeType bool) error {
	for _, i := range items {
		_, err := s.prodSrv.UpdateAProductQuantity(ctx, rollbackData, i.ID, i.Quantity, changeType)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) checkoutRollbackHandler(ctx context.Context, rollbackData map[primitive.ObjectID]interface{}) error {
	if rollbackData == nil {
		return fmt.Errorf("none data to rollback")
	}
	for key, data := range rollbackData {
		_, err := s.prodSrv.UpdateAProductQuantity(ctx, rollbackData, key, data.(*uint64), true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) saveSignature(ctx context.Context, payOSCode string, signature string) error {
	err := s.rdb.Set(ctx, payOSCode, signature, 5*time.Minute).Err()
	return err
}

// Hàm helper dùng chung cho cả Member và Guest
func (s *service) processCheckout(ctx context.Context, orderItems []model.OrderItem, deliveryAddress userModel.DeliveryAddress, promotionIDs []primitive.ObjectID, userID *primitive.ObjectID) (map[string]interface{}, map[primitive.ObjectID]interface{}, error) {

	if len(orderItems) == 0 {
		return nil, nil, fmt.Errorf("cart is empty")
	}
	if len(deliveryAddress.Phone) < 5 {
		return nil, nil, fmt.Errorf("invalid phone number")
	}

	var itemIds []primitive.ObjectID

	for _, i := range orderItems {
		itemIds = append(itemIds, i.ProductID)
	}
	items, err := s.prodSrv.GetProductByID(ctx, itemIds)
	if err != nil {
		return nil, nil, err
	}
	var rollbackData map[primitive.ObjectID]interface{}
	err = s.updateQuantityForMultiple(ctx, rollbackData, items, false)

	var totalPrice int64
	for _, item := range items {
		totalPrice += *item.Price
	}

	phoneNum := deliveryAddress.Phone
	phoneNumLastFour := phoneNum[len(phoneNum)-4:]
	suffix, _ := strconv.ParseInt(phoneNumLastFour, 10, 64)
	orderCode := time.Now().Unix()*10000 + suffix

	orderDescription := "Thanh-toan-don-hang-" + strconv.FormatInt(orderCode, 10)
	cancelURL := s.cfg.FEAddr + "/cancel"
	returnURL := s.cfg.FEAddr + "/return"

	rawSignature := fmt.Sprintf("amount=%d&cancelUrl=%s&description=%s&orderCode=%d&returnUrl=%s", totalPrice, cancelURL, orderDescription, orderCode, returnURL)
	signature := utils.ComputeHmac256(rawSignature, s.cfg.PayOSChecksum)

	order := &model.Order{
		UserID:           userID,
		OrderNumber:      orderCode,
		OrderItems:       orderItems,
		PromotionIDs:     promotionIDs,
		TotalPrice:       totalPrice,
		DeliveryAddress:  deliveryAddress,
		Status:           model.StatusPending,
		PaymentSignature: signature,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if _, err := s.repo.CreateAOrder(ctx, order); err != nil {
		return nil, rollbackData, fmt.Errorf("error when create an order: %w", err)
	}

	payOSReq := map[string]interface{}{
		"orderCode":   orderCode,
		"amount":      totalPrice,
		"description": orderDescription,
		"cancelUrl":   cancelURL,
		"returnUrl":   returnURL,
		"expiredAt":   time.Now().Add(time.Minute * 7).Unix(),
		"signature":   signature,
	}

	result, err := s.payOS.CreatePaymentLink(payOSReq)
	if err != nil {
		filter := bson.M{"order_number": orderCode}
		update := bson.M{"$set": bson.M{"status": model.StatusCancelled, "updated_at": time.Now()}}
		_ = s.repo.UpdateAOrder(ctx, filter, update)
		return nil, rollbackData, fmt.Errorf("error when create a payment link: %w", err)
	}

	err = s.saveSignature(ctx, result["code"].(string), result["signature"].(string))
	if err != nil {
		return nil, rollbackData, fmt.Errorf("error when save payment signature: %w", err)
	}

	delete(result, "checkoutUrl")
	delete(result, "paymentLinkId")
	return result, nil, nil
}

func (s *service) MemberCheckout(ctx context.Context, req model.MemberOrderRequest, userIdStr string) (data map[string]interface{}, err error) {
	var rollbackData map[primitive.ObjectID]interface{}
	defer func() {
		if r := recover(); (r != nil || err != nil) && len(rollbackData) > 0 {
			fmt.Println("Err caught in MemberCheckout")
			_ = s.checkoutRollbackHandler(ctx, rollbackData)
			return
		}
	}()

	var userId primitive.ObjectID
	userId, err = utils.ConvertStringToID(userIdStr)
	if err != nil {
		return nil, err
	}
	data, rollbackData, err = s.processCheckout(ctx, req.OrderItems, req.DeliveryAddress, req.PromotionIDs, &userId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (s *service) GuestCheckout(ctx context.Context, req model.GuestOrderRequest) (data map[string]interface{}, err error) {
	var rollbackData map[primitive.ObjectID]interface{}
	defer func() {
		if r := recover(); (r != nil || err != nil) && len(rollbackData) > 0 {
			fmt.Println("Err caught in MemberCheckout")
			_ = s.checkoutRollbackHandler(ctx, rollbackData)
			return
		}
	}()
	data, rollbackData, err = s.processCheckout(ctx, req.OrderItems, req.DeliveryAddress, req.PromotionIDs, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) UpdateOrderStatus(ctx context.Context, orderCode int, status string) error {
	filter := bson.M{"order_number": orderCode}
	update := bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}}
	err := s.repo.UpdateAOrder(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
