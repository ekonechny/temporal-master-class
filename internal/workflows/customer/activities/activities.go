package activities

import (
	"temporal-master-class/internal/services/assortment"
)

type Activities struct {
	assortmentClient *assortment.Client
}

func Register(ac *assortment.Client) *Activities {
	return &Activities{
		assortmentClient: ac,
	}
}
