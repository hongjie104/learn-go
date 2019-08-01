package v1

import (
	"github.com/gin-gonic/gin"

	response "../../../pkg/app"
	"../../../pkg/e"
	"../../../pkg/utils"
)

// GetUser a
func GetUser(c *gin.Context) {
	response := response.Gin{C: c}
	response.Success(gin.H{
		"name":   "鸿杰",
		"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"userid": 3630,
	})
}

// LoginData 绑定为json
type LoginData struct {
	User     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login a
func Login(c *gin.Context) {
	response := response.Gin{C: c}
	var loginData LoginData
	if err := c.BindJSON(&loginData); err == nil {
		if loginData.User == "admin" && loginData.Password == "123" {
			token, _ := utils.GenerateToken(loginData.User, loginData.Password)
			response.Success(gin.H{"token": token})
		} else {
			response.Fail(e.LoginError)
		}
	} else {
		response.Fail(e.InvalidParams)
	}
}
