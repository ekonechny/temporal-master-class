package activities

import (
	"temporal-master-class/internal/services/payment"
	"temporal-master-class/internal/services/vendors"
)

type Activities struct {
	paymentClient *payment.Client
	vendorsClient *vendors.Client
}

func Register(pc *payment.Client, vc *vendors.Client) *Activities {
	return &Activities{
		paymentClient: pc,
		vendorsClient: vc,
	}
}
