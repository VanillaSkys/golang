package service

import (
	"github.com/VanillaSkys/golang/database/postgresql"
	"github.com/VanillaSkys/golang/model"
	"github.com/VanillaSkys/golang/repository"
	"github.com/VanillaSkys/golang/request"
	"github.com/VanillaSkys/golang/response"
)

type ProductService interface {
	FindAllProduct() ([]response.GetProductsResponse, error)
	FindProductByName(name request.FindByName) (response.GetProductsResponse, error)
	CreateProduct(product request.CreateProduct) error
	UpdateProductById(id request.UpdateById) error
}

func (service Service) FindAllProduct() ([]response.GetProductsResponse, error) {
	var products []response.GetProductsResponse
	repoObj := repository.New(service.RequestId, postgresql.DB)
	result, err := repoObj.FindAll()
	if err != nil {
		return nil, err
	}
	for _, value := range result {
		product := response.GetProductsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		products = append(products, product)
	}
	return products, nil
}

func (service Service) FindProductByName(request request.FindByName) (response.GetProductsResponse, error) {
	var product response.GetProductsResponse
	repoObj := repository.New(service.RequestId, postgresql.DB)
	result, err := repoObj.FindByName(request.Name)
	if err != nil {
		return product, err
	}
	product.Id = result.Id
	product.Name = result.Name
	return product, nil
}

func (service Service) CreateProduct(product request.CreateProduct) error {
	repoObj := repository.New(service.RequestId, postgresql.DB)
	productModel := model.Product{
		Id:   product.Id,
		Name: product.Name,
	}
	err := repoObj.Save(productModel)
	if err != nil {
		return err
	}
	return nil
}

func (service Service) UpdateProductById(request request.UpdateById) error {
	repoObj := repository.New(service.RequestId, postgresql.DB)
	err := repoObj.Update(request.Id, request.Name)
	if err != nil {
		return err
	}
	return nil
}

func (service Service) DeleteById(request request.DeleteById) error {
	repoObj := repository.New(service.RequestId, postgresql.DB)
	err := repoObj.Delete(request.Id)
	if err != nil {
		return err
	}
	return nil
}
