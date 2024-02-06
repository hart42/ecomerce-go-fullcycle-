package service

import (
	"github.com/hart42/ecomerce-go-fullcycle-/internal/database"
	"github.com/hart42/ecomerce-go-fullcycle-/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDb database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: categoryDb}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
