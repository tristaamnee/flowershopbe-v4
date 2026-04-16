package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type PayOSProvider struct {
	PayOSLink string
	ClientID  string
	ApiKey    string
	Checksum  string
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

type PayOSPaymentRequest struct {
	OrderCode        int      `json:"order_code"`
	Amount           int      `json:"amount"`
	Description      string   `json:"description"`
	BuyerName        *string  `json:"buyer_name"`
	BuyerCompanyName *string  `json:"buyer_company_name"`
	BuyerTaxCode     *string  `json:"buyer_tax_code"`
	BuyerAddress     *string  `json:"buyer_address"`
	BuyerEmail       *string  `json:"buyer_email"`
	BuyerPhone       *string  `json:"buyer_phone"`
	Items            *[]Item  `json:"items"`
	CancelURL        string   `json:"cancel_url"`
	ReturnURL        string   `json:"return_url"`
	Invoice          *Invoice `json:"invoice"`
	ExpireAt         string   `json:"expire_at"`
	Signature        string   `json:"signature"`
}

func NewPayOSProvider() *PayOSProvider {
	return &PayOSProvider{
		PayOSLink: os.Getenv("PAYOS"),
		ClientID:  os.Getenv("PAYOS_CLIENT_ID"),
		ApiKey:    os.Getenv("PAYOS_API_KEY"),
		Checksum:  os.Getenv("PAYOS_CHECKSUM"),
	}
}

func (p *PayOSProvider) CreatePaymentLink(payload interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(payload)
	PaymentCreateLink := p.PayOSLink + "/v2/payment-request"
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", PaymentCreateLink, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.ApiKey)
	req.Header.Set("x-client-id", p.ClientID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
