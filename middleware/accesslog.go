package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wangming1993/gohttp/log"
)

func AccessLog(c *gin.Context) {
	log.Write(map[string]interface{}{
		"time": time.Now(),
		"path": c.Request.RequestURI,
	})

	c.Next()

}
