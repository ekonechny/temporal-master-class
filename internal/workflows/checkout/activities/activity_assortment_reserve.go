package activities

import (
	"context"

	"temporal-master-class/internal/generated/temporal"
)

func (a *Activities) AssortmentReserve(ctx context.Context, req *temporal.AssortmentReserveRequest) error {
	return a.assortmentClient.Reserve(ctx, req.Products)
}
