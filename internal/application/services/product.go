package services

import "context"

type Product struct {
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) CreateProduct(ctx context.Context) {

}
