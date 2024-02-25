package controller

import (
	"github.com/VanillaSkys/golang/request"
	"github.com/VanillaSkys/golang/service"
)

// func (controller Controller) GetProduct() (response.GetProductsResponse, error) {

// }

func (controller Controller) CreateProduct(request request.CreateProduct) error {
	serviceObj := service.New(controller.RequestId)
	err := serviceObj.CreateProduct(request.Id, request.Name)
	if err != nil {
		return err
	}
	return nil
}
