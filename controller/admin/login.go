package admin

import (
	"MiShop/logger"
	"MiShop/logic"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginControllerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}
func LoginControllerCaptcha(c *gin.Context) {
	id, s, err := logic.CaptMake()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": s,
	})
}
func DoLoginController(c *gin.Context) {
	captchID := c.PostForm("captchaId")
	vertifyValue := c.PostForm("verifyValue")
	username := c.PostForm("username")
	password := c.PostForm("password")
	pass := logic.Md5(password)
	if flag := logic.VertifyCaptcha(captchID, vertifyValue); flag == true {
		logic.Login(username, pass, c)
	} else {
		logic.ErrorReply(c, "验证失败", "/admin/login")
	}
}
func LogOutController(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userInfo")
	session.Save()
	logic.SuccessReply(c, "登出成功", "/admin/login")
}
