package controller

import (
	"github.com/gin-gonic/gin"
)

// Home is homepage
func Home(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"status": "OK",
	})
}
