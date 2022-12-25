package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessReply(c *gin.Context, msg string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     msg,
		"redirectUrl": redirectUrl,
	})
}
func ErrorReply(c *gin.Context, msg string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     msg,
		"redirectUrl": redirectUrl,
	})
}
