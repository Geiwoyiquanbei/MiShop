package routers

import (
	"MiShop/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	AdminGroups := r.Group("/admin")
	AdminGroups.GET("login", admin.LoginControllerIndex)
	AdminGroups.GET("captcha", admin.LoginControllerCaptcha)
	AdminGroups.POST("doLogin", admin.DoLoginController)

	AdminGroups.GET("manager", admin.ManagerController)
	AdminGroups.GET("manager/add", admin.ManagerAddController)
	AdminGroups.GET("manager/edit", admin.ManagerEditController)
	AdminGroups.GET("manager/delete", admin.ManagerDeleteController)

	AdminGroups.GET("focus", admin.FocusController)
	AdminGroups.GET("focus/add", admin.FocusAddController)
	AdminGroups.GET("focus/edit", admin.FocusEditController)
	AdminGroups.GET("focus/delete", admin.FocusDeleteController)
}
