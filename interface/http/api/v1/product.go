package v1

import (
	"fmt"

	"github.com/VanillaSkys/golang/controller"
	"github.com/VanillaSkys/golang/request"
	"github.com/VanillaSkys/golang/response"
	"github.com/gin-gonic/gin"
)

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

func init() {
	methodRoutes[ROUTE_POST]["/product"] = CreateProduct
}
