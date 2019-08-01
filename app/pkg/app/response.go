package response

import (
	"net/http"

	"../e"
	"github.com/gin-gonic/gin"
)

// Gin a
type Gin struct {
	C *gin.Context
}

// Success a
func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Fail a
func (g *Gin) Fail(errCode int) {
	g.C.JSON(http.StatusOK, gin.H{
		"success": false,
		"errCode": errCode,
		"errMsg":  e.GetMsg(errCode),
	})
}
