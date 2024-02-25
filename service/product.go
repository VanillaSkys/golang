package service

import (
	"github.com/VanillaSkys/golang/database/postgresql"
	"github.com/VanillaSkys/golang/repository"
)

type ProductService interface {
	CreateProduct(name string) error
}

func (service Service) CreateProduct(id string, name string) error {
	repoObj := repository.New(service.RequestId, postgresql.DB)
	err := repoObj.Save(id, name)
	if err != nil {
		return err
	}
	return nil
}
