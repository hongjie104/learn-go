package v1

import (
	"fmt"
	"net/http"

	"../../../models"
	"github.com/gin-gonic/gin"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	tag := models.GetTags(1, 10)
	fmt.Println(tag)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tag})
	// c.JSON(http.StatusOK, gin.H{"data": tag, "success": true})
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
}
