package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ManagerController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})
}
func ManagerAddController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{})
}
func ManagerEditController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}
func ManagerDeleteController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "successful",
	})
}
