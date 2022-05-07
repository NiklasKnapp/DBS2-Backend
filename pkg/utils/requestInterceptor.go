package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

//not in use
func RequestInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[REQUEST INTERCEPTOR]:", c.Request.Method, c.Request.URL, c.Request.ContentLength, c.Request.Header, c.Request.Body)
		c.Next()
	}
}
