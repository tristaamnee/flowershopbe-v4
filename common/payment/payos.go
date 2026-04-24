package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tristaamne/flowershopbe-v4/common/config"
)

type payOSProvider struct {
	// chinh lai may cai anh huong
	cfg    *config.Config
	client *http.Client
	rdb    *redis.Client
}

func NewPayOSProvider(cfg *config.Config, rdb *redis.Client) PaymentProvider {
	return &payOSProvider{
		cfg: cfg,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		rdb: rdb,
	}
}

type Item struct {
	Name          string
	Quantity      int
	Price         float64
	Unit          string
	TaxPercentage float64
}

type Invoice struct {
	BuyerNotGetInvoice bool
	TaxPercentage      float64
}

type PayOSWebhookBody struct {
	Code    string `json:"code"`
	Desc    string `json:"desc"`
	Success bool   `json:"success"`
	Data    struct {
		OrderCode              int     `json:"orderCode"`
		Amount                 int     `json:"amount"`
		Description            string  `json:"description"`
		AccountNumber          string  `json:"accountNumber"`
		Reference              string  `json:"reference"`
		TransactionDateTime    string  `json:"transactionDateTime"`
		Currency               string  `json:"currency"`
		PaymentLinkId          string  `json:"paymentLinkId"`
		Code                   string  `json:"code"`
		Desc                   string  `json:"desc"`
		CounterAccountBankId   *string `json:"counterAccountBankId"`
		CounterAccountBankName *string `json:"counterAccountBankName"`
		CounterAccountName     *string `json:"counterAccountName"`
		CounterAccountNumber   *string `json:"counterAccountNumber"`
		VirtualAccountName     *string `json:"virtualAccountName"`
		VirtualAccountNumber   *string `json:"virtualAccountNumber"`
	}
	Signature string `json:"signature"`
}

type CancelPaymentRequest struct {
	CancellationReason string `json:"cancellation_reason"`
}

func (p *payOSProvider) CheckWebhookSignature(ctx context.Context, body PayOSWebhookBody) error {
	if body.Success {
		orderCode := body.Code
		storedSignature, err := p.rdb.Get(ctx, orderCode).Result()
		if err != nil {
			return err
		}
		if storedSignature != body.Signature {
			return errors.New("invalid signature")
		}
		return nil
	}

	return errors.New("payment not success")

}

func (p *payOSProvider) CreatePaymentLink(payload interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	PaymentCreateLink := p.cfg.PayOS + "/v2/payment-request"
	req, err := http.NewRequest("POST", PaymentCreateLink, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.cfg.PayOSApiKey)
	req.Header.Set("x-client-id", p.cfg.PayOSClientID)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (p *payOSProvider) CancelPaymentLink(paymentLinkId, cancelReason string) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(CancelPaymentRequest{CancellationReason: cancelReason})
	if err != nil {
		return nil, err
	}

	PaymentCancelLink := fmt.Sprintf("%s/v2/payment-requests/%s/cancel", p.cfg.PayOS, paymentLinkId)
	req, err := http.NewRequest("POST", PaymentCancelLink, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.cfg.PayOSApiKey)
	req.Header.Set("x-client-id", p.cfg.PayOSClientID)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
