package assortment

import (
	"context"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"

	"temporal-master-class/internal/generated/temporal"
)

type Client struct {
	faker *gofakeit.Faker
	http.Client
}

func (c *Client) GetProducts(ctx context.Context, productIDs []string) ([]*temporal.AssortmentProduct, error) {
	// TODO: на самом деле здесь должен быть сгенеренный http-клиент (или просто http-клиент)
	// но мне пока его лень писать, поэтому тут будет мок
	res := make([]*temporal.AssortmentProduct, len(productIDs))
	for i := range productIDs {
		res[i] = &temporal.AssortmentProduct{
			Id:     productIDs[i],
			Name:   c.faker.ProductName(),
			Price:  int32(c.faker.Price(100, 9999)),
			Stocks: int32(c.faker.IntRange(5, 32)),
		}
	}
	return res, nil
}

func (c *Client) Reserve(ctx context.Context, products []*temporal.AssortmentReserveProductRequest) error {
	// TODO: на самом деле здесь должен быть сгенеренный http-клиент (или просто http-клиент)
	// но мне пока его лень писать, поэтому тут будет мок
	return nil
}

func New() *Client {
	return &Client{
		faker: gofakeit.New(0),
	}
}
