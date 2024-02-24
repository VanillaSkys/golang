package service

import (
	"github.com/VanillaSkys/golang/repository"
)

type ProductService interface {
	CreateProduct(name string) error
}

func (service Service) CreateProduct(name string) error {
	repoObj := repository.New(service.RequestId)
	if err := repoObj.Save(name); err != nil {
		return err
	}
	return nil
}
