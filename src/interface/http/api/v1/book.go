package v1

import (
	"github.com/VanillaSkys/golang/controller"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	requestID, _ := c.Get("request_id")
	con := controller.New(requestID.(string))
	out := con.GetBook()
	c.JSON(200, gin.H{
		"message":   out,
		"requestID": requestID,
	})
}

func CreateBook(c *gin.Context) {
	input := controller.CreateBookInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(507)
		return
	}
	requestID, _ := c.Get("request_id")
	con := controller.New(requestID.(string))
	out := con.CreateBook(input)
	c.JSON(200, out)
}

func init() {
	methodRoutes[ROUTE_GET]["/books"] = GetBooks
	methodRoutes[ROUTE_POST]["/books"] = CreateBook
}
