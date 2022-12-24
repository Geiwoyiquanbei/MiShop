package admin

import (
	"MiShop/logger"
	"MiShop/logic"
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
	if flag := logic.VertifyCaptcha(captchID, vertifyValue); flag == true {
		c.String(http.StatusOK, "验证成功")
	} else {
		c.String(http.StatusOK, "验证失败")
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "successful",
	})
}
