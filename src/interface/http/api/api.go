package api

import (
	v1 "github.com/VanillaSkys/golang/interface/http/api/v1"
	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.Engine) {
	apiGroup := router.Group("/api")
	v1.AddRoutes(apiGroup)
}
