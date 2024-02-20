package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exist := c.Get("request_id"); !exist {
			c.Set("request_id", uuid.New().String())
		}
		requestId := uuid.New().String()
		c.Set("request_id", requestId)
		c.Next()
	}
}
