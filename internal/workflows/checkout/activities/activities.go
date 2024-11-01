package activities

import (
	"context"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/services/assortment"
	"temporal-master-class/internal/services/payment"
)

type Activities struct {
	assortmentClient *assortment.Client
	paymentClient    *payment.Client
}

func (a *Activities) PaymentCancel(ctx context.Context, req *temporal.PaymentCancelRequest) error {
	//TODO implement me
	panic("implement me")
}

func Register(ac *assortment.Client, pc *payment.Client) *Activities {
	return &Activities{
		assortmentClient: ac,
		paymentClient:    pc,
	}
}
