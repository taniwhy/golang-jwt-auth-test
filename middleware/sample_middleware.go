package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
