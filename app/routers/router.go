package routers

import (
	"net/http"

	"../pkg/setting"
	v1 "./api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter a
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "success": true})
	})

	apiV1 := r.Group("api/v1")
	// apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.GET("/test", v1.Test)
	}

	return r
}
