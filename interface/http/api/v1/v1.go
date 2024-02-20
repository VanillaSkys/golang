package v1

import (
	"github.com/gin-gonic/gin"
)

const (
	ROUTE_GET    = "GET"
	ROUTE_POST   = "POST"
	ROUTE_PUT    = "PUT"
	ROUTE_DELETE = "DELETE"
)

var methodRoutes = map[string]map[string]gin.HandlerFunc{
	ROUTE_GET:    make(map[string]gin.HandlerFunc),
	ROUTE_POST:   make(map[string]gin.HandlerFunc),
	ROUTE_PUT:    make(map[string]gin.HandlerFunc),
	ROUTE_DELETE: make(map[string]gin.HandlerFunc),
}

func AddRoutes(router *gin.RouterGroup) {
	v1Router := router.Group("/v1")

	for method, routes := range methodRoutes {
		if method == ROUTE_GET {
			for routeName, routeFunc := range routes {
				v1Router.GET(routeName, routeFunc)
			}
		} else if method == ROUTE_POST {
			for routeName, routeFunc := range routes {
				v1Router.POST(routeName, routeFunc)
			}
		} else if method == ROUTE_PUT {
			for routeName, routeFunc := range routes {
				v1Router.PUT(routeName, routeFunc)
			}
		} else if method == ROUTE_DELETE {
			for routeName, routeFunc := range routes {
				v1Router.DELETE(routeName, routeFunc)
			}
		}
	}
}
