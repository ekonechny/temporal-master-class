package activities

import (
	"context"

	"temporal-master-class/internal/generated/temporal"
)

func (a *Activities) GetVendorOrder(ctx context.Context, req *temporal.VendorOrderRequest) (*temporal.VendorOrderResponse, error) {
	return a.vendorsClient.GetOrder(ctx, req.Id)
}
