package customer

import (
	"errors"

	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
)

func (w *Workflow) UpdateCart(ctx workflow.Context, request *temporal.UpdateCartRequest) (*temporal.Cart, error) {
	// Формируем список продуктов, которые надо получить из ассортимента
	productIDs := make([]string, len(request.Products))
	for i := range request.Products {
		productIDs[i] = request.Products[i].GetId()
	}

	// Получаем те самые продукты c помощью activity
	assortmentGetProductsReq := &temporal.AssortmentGetProductsRequest{
		Ids: productIDs,
	}
	assortment, err := temporal.AssortmentGetProducts(ctx, assortmentGetProductsReq)
	if err != nil {
		return nil, err
	}

	// Проверяем, что все продукты есть и их хватает для добавления в корзину
	stocks := make(map[string]*temporal.AssortmentProduct)
	for _, p := range assortment.Products {
		stocks[p.Id] = p
	}
	products := make([]*temporal.Product, 0, len(request.Products))
	for _, p := range request.Products {
		assortmentProduct, ok := stocks[p.Id]
		if !ok {
			return nil, errors.New("the product is not in stock")
		}
		if assortmentProduct.Stocks < p.Qty {
			return nil, errors.New("not enough goods in stock")
		}
		products = append(products, &temporal.Product{
			Id:    assortmentProduct.Id,
			Name:  assortmentProduct.Name,
			Price: assortmentProduct.Price,
			Inn:   assortmentProduct.Inn,
			Qty:   p.Qty,
		})
	}
	total := calculateTotal(products)

	w.cart = &temporal.Cart{
		// Добавляем сюда генерацию UUID
		//Id:       uuid.NewString(),
		Products: products,
		Total:    total,
	}

	// Перезапускаем сервис видим другой uuid

	/*
		// Добавляем через SideEffect:
		// https://docs.temporal.io/develop/go/side-effects
		encodedValue := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return uuid.NewString()
		})
		w.cart = &temporal.Cart{
			Products: products,
			Total:    total,
		}
		if err := encodedValue.Get(&w.cart.Id); err != nil {
			return nil, err
		}
	*/

	// Опачки
	// wID customers/c527a4b1-10e0-4b0b-a555-8dd49f28055b RunID 81f1229f-a662-42b8-896f-af5d711d836b Attempt 1 Error [TMPRL1100] No cached result found for side effectID=1. KnownSideEffects=[] StackTrace coroutine temporal.Customer.UpdateCart [panic]:

	/*if w.cart == nil {
		w.cart = &temporal.Cart{}
	}
	w.cart.Products = products
	w.cart.Total = total
	// Если уже идентификатор задан, то скипаем
	if w.cart.Id != "" {
		return w.cart, nil
	}
	// Берем динамически создание версии
	versionID := workflow.GetCurrentUpdateInfo(ctx).ID
	v := workflow.GetVersion(ctx, fmt.Sprintf("cartUpdate-%s", versionID), workflow.DefaultVersion, 1)
	if v > 0 {
		encodedValue := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
			return uuid.NewString()
		})

		if err := encodedValue.Get(&w.cart.Id); err != nil {
			return nil, err
		}
	}*/
	return w.cart, nil
}

func calculateTotal(products []*temporal.Product) int32 {
	var total int32
	for i := range products {
		total += products[i].Qty * products[i].Price
	}
	return total
}
