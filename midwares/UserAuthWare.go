package midwares

import (
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-gonic/gin"
)

func UserAuthMidWare(c *gin.Context) {
	user := models.User{}
	ok := logic.Cookie.Get(c, "userinfo", &user)
	if !ok || len(user.Phone) != 11 {
		{
			c.Redirect(302, "/pass/login")
		}
	}
}
