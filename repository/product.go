package repository

import (
	"errors"

	"github.com/VanillaSkys/golang/model"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByName(name string) (model.Product, error)
	Save(product model.Product) error
	Update(id string, name string) error
}

func (repository Repository) FindAll() ([]model.Product, error) {
	var product []model.Product
	result := repository.DB.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repository Repository) FindByName(name string) (model.Product, error) {
	var product model.Product
	result := repository.DB.Find(&product, name)
	if result != nil {
		return product, nil
	}
	return product, errors.New("product is not found")
}

func (repository Repository) Save(product model.Product) error {
	result := repository.DB.Save(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository Repository) Update(id string, name string) error {
	result := repository.DB.Model(&model.Product{}).Where("id = ?", id).Update("name", name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
