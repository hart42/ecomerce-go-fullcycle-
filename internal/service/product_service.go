package service

import (
	"github.com/hart42/ecomerce-go-fullcycle-/internal/database"
	"github.com/hart42/ecomerce-go-fullcycle-/internal/entity"
)

type ProductService struct {
	ProductDb database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDb: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDb.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
