package vendors

import (
	"context"

	"github.com/google/uuid"

	"temporal-master-class/internal/generated/temporal"
)

type Client struct{}

func (c *Client) CreateOrder(ctx context.Context, req *temporal.CreateVendorOrderRequest) (*temporal.CreateVendorOrderResponse, error) {
	return &temporal.CreateVendorOrderResponse{
		Id:     uuid.NewString(),
		Status: temporal.VendorOrderStatus_VendorOrderNew,
	}, nil
}

func (c *Client) GetOrder(ctx context.Context, orderID string) (*temporal.VendorOrderResponse, error) {
	return &temporal.VendorOrderResponse{
		Id:     orderID,
		Status: temporal.VendorOrderStatus_VendorOrderNew,
	}, nil
}

func New() *Client {
	return &Client{}
}
