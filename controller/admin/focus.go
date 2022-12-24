package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FocusController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}
func FocusAddController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func FocusEditController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}
func FocusDeleteController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "successful",
	})
}
