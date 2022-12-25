package logic

import (
	"MiShop/dao/mysql"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(username, pass string, c *gin.Context) {
	userInfo := mysql.Login(username, pass)
	if len(userInfo) > 0 {
		session := sessions.Default(c)
		marshal, _ := json.Marshal(userInfo)
		session.Set("userInfo", string(marshal))
		session.Save()
		SuccessReply(c, "登录成功", "/admin")
	} else {
		ErrorReply(c, "用户名或密码错误", "/admin/login")
	}
}
