package controller

import (
	"github.com/VanillaSkys/golang/request"
	"github.com/VanillaSkys/golang/response"
	"github.com/VanillaSkys/golang/service"
)

func (controller Controller) FindAllProduct() ([]response.GetProductsResponse, error) {
	serviceObj := service.New(controller.RequestId)
	response, err := serviceObj.FindAllProduct()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (controller Controller) CreateProduct(request request.CreateProduct) error {
	serviceObj := service.New(controller.RequestId)
	err := serviceObj.CreateProduct(request)
	if err != nil {
		return err
	}
	return nil
}
