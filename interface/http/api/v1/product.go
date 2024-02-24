package v1

import (
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
	if err != nil {
		responseErr := response.ResponseError{}
		responseErr.Status = 507
		responseErr.Error = err
	}
	responseData.Status = 200
	c.JSON(200, responseData)
}

func init() {
	methodRoutes[ROUTE_POST]["/product"] = CreateBook
}
