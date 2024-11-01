package payment

import (
	"context"

	"github.com/google/uuid"

	"temporal-master-class/internal/generated/temporal"
)

type Client struct{}

func (c *Client) CreatePayment(ctx context.Context, req *temporal.CreatePaymentRequest) (*temporal.CreatePaymentResponse, error) {
	// TODO: на самом деле здесь должен быть сгенеренный http-клиент (или просто http-клиент)
	// но мне пока его лень писать, поэтому тут будет мок
	return &temporal.CreatePaymentResponse{
		Id:     uuid.NewString(),
		Status: temporal.PaymentStatus_PaymentStatusNew,
	}, nil
}

func (c *Client) GetPayment(ctx context.Context, paymentId string) (*temporal.PaymentStatusResponse, error) {
	// TODO: на самом деле здесь должен быть сгенеренный http-клиент (или просто http-клиент)
	// но мне пока его лень писать, поэтому тут будет мок
	return &temporal.PaymentStatusResponse{
		Id:     paymentId,
		Status: temporal.PaymentStatus_PaymentStatusNew,
	}, nil
}

func New() *Client {
	return &Client{}
}
