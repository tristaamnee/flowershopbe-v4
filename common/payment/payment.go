package payment

import "context"

type PaymentProvider interface {
	CreatePaymentLink(payload interface{}) (map[string]interface{}, error)
	CancelPaymentLink(paymentLinkId, cancelReason string) (map[string]interface{}, error)
	CheckWebhookSignature(ctx context.Context, body PayOSWebhookBody) error
}
