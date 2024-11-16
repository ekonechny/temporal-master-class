package activities

import (
	"temporal-master-class/internal/services/assortment"
	"temporal-master-class/internal/services/payment"
)

type Activities struct {
	assortmentClient *assortment.Client
	paymentClient    *payment.Client
}

func Register(ac *assortment.Client, pc *payment.Client) *Activities {
	return &Activities{
		assortmentClient: ac,
		paymentClient:    pc,
	}
}
