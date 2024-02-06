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

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDb.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	_, err := ps.ProductDb.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, err
}
