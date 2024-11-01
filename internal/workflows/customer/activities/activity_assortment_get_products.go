package activities

import (
	"context"

	"temporal-master-class/internal/generated/temporal"
)

func (a *Activities) AssortmentGetProducts(ctx context.Context, req *temporal.AssortmentGetProductsRequest) (*temporal.AssortmentGetProductsResponse, error) {
	p, err := a.assortmentClient.GetProducts(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return &temporal.AssortmentGetProductsResponse{
		Products: p,
	}, nil
}
