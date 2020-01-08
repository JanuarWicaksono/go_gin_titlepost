package handler

import (
	"github.com/gin-gonic/gin"
)

func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		// results := feed.GetAll()
		// c.JSON(http.StatusOK, results)
	}
}
