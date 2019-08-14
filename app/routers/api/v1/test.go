package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../../../pkg/log"
	"../../../pkg/utils"
)

// Test a
func Test(c *gin.Context) {
	// token, _ := utils.GenerateToken("tom", "123")
	// c.JSON(http.StatusOK, gin.H{"test": true, "data": token})
	t, err := utils.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRvbSIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNTY0MjI3NzQ1LCJpc3MiOiJsZWFybi1nbyJ9.67I3ccNeCPRheTc-YoUnPeZwTXNb6u2d8dLIqktAQT0")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"test": true, "data": t.Password})
	} else {
		log.LogError(err)
		c.JSON(http.StatusOK, gin.H{"success": false, "data": nil})
	}
}
