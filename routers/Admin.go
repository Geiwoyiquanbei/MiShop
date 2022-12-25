package routers

import (
	"MiShop/controller/admin"
	"MiShop/midwares"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	AdminGroups := r.Group("/admin", midwares.InitAdminAuthMiddleware)
	AdminGroups.GET("/", admin.MainController)
	AdminGroups.GET("/welcome", admin.WelcomeController)

	AdminGroups.GET("/login", admin.LoginControllerIndex)
	AdminGroups.GET("/captcha", admin.LoginControllerCaptcha)
	AdminGroups.GET("/loginOut", admin.LogOutController)
	AdminGroups.POST("/doLogin", admin.DoLoginController)

	AdminGroups.GET("/manager", admin.ManagerController)
	AdminGroups.GET("/manager/add", admin.ManagerAddController)
	AdminGroups.POST("/manager/doAdd", admin.ManagerDoAddController)
	AdminGroups.GET("/manager/edit", admin.ManagerEditController)
	AdminGroups.POST("/manager/doEdit", admin.ManagerDoEditController)
	AdminGroups.GET("/manager/delete", admin.ManagerDeleteController)

	AdminGroups.GET("/focus", admin.FocusController)
	AdminGroups.GET("/focus/add", admin.FocusAddController)
	AdminGroups.GET("/focus/edit", admin.FocusEditController)
	AdminGroups.GET("/focus/delete", admin.FocusDeleteController)

	AdminGroups.GET("/role", admin.RoleController)
	AdminGroups.GET("/role/add", admin.RoleAddController)
	AdminGroups.POST("/role/doAdd", admin.RoleDoAddController)
	AdminGroups.GET("/role/edit", admin.RoleEditController)
	AdminGroups.POST("/role/doEdit", admin.RoleDoEditController)
	AdminGroups.GET("/role/delete", admin.RoleDeleteController)

}
