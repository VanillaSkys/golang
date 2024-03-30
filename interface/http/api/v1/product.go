package v1

import (
	"fmt"

	"github.com/VanillaSkys/golang/controller"
	"github.com/VanillaSkys/golang/request"
	"github.com/VanillaSkys/golang/response"
	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	responseData := response.ResponseData{}

	requestID, _ := c.Get("request_id")
	controllerObj := controller.New(requestID.(string))
	output, err := controllerObj.FindAllProduct()
	fmt.Println(err)
	if err != nil {
		responseErr := response.ResponseError{}
		responseErr.Status = 404
		responseErr.Error = err
	}
	responseData.Status = 200
	responseData.Data = output
	c.JSON(200, responseData)
}

func CreateProduct(c *gin.Context) {
	input := request.CreateProduct{}
	responseData := response.ResponseData{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(507)
		return
	}
	requestID, _ := c.Get("request_id")
	controllerObj := controller.New(requestID.(string))
	err := controllerObj.CreateProduct(input)
	fmt.Println(err)
	if err != nil {
		responseErr := response.ResponseError{}
		responseErr.Status = 404
		responseErr.Error = err
	}
	responseData.Status = 200
	responseData.Data = input
	c.JSON(200, responseData)
}

func UpdateProduct(c *gin.Context) {
	input := request.UpdateById{}
	responseData := response.ResponseData{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(507)
		return
	}
	requestID, _ := c.Get("request_id")
	controllerObj := controller.New(requestID.(string))
	err := controllerObj.UpdateProduct(input)
	fmt.Println(err)
	if err != nil {
		responseErr := response.ResponseError{}
		responseErr.Status = 404
		responseErr.Error = err
	}
	responseData.Status = 200
	responseData.Data = input
	c.JSON(200, responseData)
}

func DeleteProduct(c *gin.Context) {
	input := request.DeleteById{}
	responseData := response.ResponseData{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(507)
		return
	}
	requestID, _ := c.Get("request_id")
	controllerObj := controller.New(requestID.(string))
	err := controllerObj.DeleteProduct(input)
	fmt.Println(err)
	if err != nil {
		responseErr := response.ResponseError{}
		responseErr.Status = 404
		responseErr.Error = err
	}
	responseData.Status = 200
	responseData.Data = input
	c.JSON(200, responseData)
}

func init() {
	methodRoutes[ROUTE_GET]["/product"] = GetAllProduct
	methodRoutes[ROUTE_POST]["/product"] = CreateProduct
	methodRoutes[ROUTE_PUT]["/product"] = UpdateProduct
	methodRoutes[ROUTE_DELETE]["/product"] = DeleteProduct
}
