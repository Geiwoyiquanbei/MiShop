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
	AdminGroups.GET("/changeStatus", admin.ChangeStatusController)
	AdminGroups.GET("/changeNum", admin.ChangeNumController)
	AdminGroups.GET("/FlushAll", admin.FlushAll)

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
	AdminGroups.POST("/focus/doAdd", admin.FocusDoAddController)
	AdminGroups.GET("/focus/edit", admin.FocusEditController)
	AdminGroups.POST("/focus/doEdit", admin.FocusDoEditController)
	AdminGroups.GET("/focus/delete", admin.FocusDeleteController)

	AdminGroups.GET("/role", admin.RoleController)
	AdminGroups.GET("/role/add", admin.RoleAddController)
	AdminGroups.POST("/role/doAdd", admin.RoleDoAddController)
	AdminGroups.GET("/role/edit", admin.RoleEditController)
	AdminGroups.POST("/role/doEdit", admin.RoleDoEditController)
	AdminGroups.GET("/role/delete", admin.RoleDeleteController)
	AdminGroups.GET("/role/auth", admin.RoleAuthController)
	AdminGroups.POST("/role/doAuth", admin.RoleDoAuthController)

	AdminGroups.GET("/access", admin.AccessController)
	AdminGroups.GET("/access/add", admin.AccessAddController)
	AdminGroups.POST("/access/doAdd", admin.AccessDoAddController)
	AdminGroups.GET("/access/edit", admin.AccessEditController)
	AdminGroups.POST("/access/doEdit", admin.AccessDoEditController)
	AdminGroups.GET("/access/delete", admin.AccessDeleteController)

	AdminGroups.GET("/goodsCate", admin.GoodsCateController)
	AdminGroups.GET("/goodsCate/add", admin.GoodsCateAddController)
	AdminGroups.POST("/goodsCate/doAdd", admin.GoodsCateDoAddController)
	AdminGroups.GET("/goodsCate/edit", admin.GoodsCateEditController)
	AdminGroups.POST("/goodsCate/doEdit", admin.GoodsCateDoEditController)
	AdminGroups.GET("/goodsCate/delete", admin.GoodsCateDeleteController)

	AdminGroups.GET("/goodsType", admin.GoodsTypeController)
	AdminGroups.GET("/goodsType/add", admin.GoodsTypeAddController)
	AdminGroups.POST("/goodsType/doAdd", admin.GoodsTypeDoAddController)
	AdminGroups.GET("/goodsType/edit", admin.GoodsTypeEditController)
	AdminGroups.POST("/goodsType/doEdit", admin.GoodsTypeDoEditController)
	AdminGroups.GET("/goodsType/delete", admin.GoodsTypeDeleteController)

	AdminGroups.GET("/goodsTypeAttribute", admin.GoodsTypeAttributeController)
	AdminGroups.GET("/goodsTypeAttribute/add", admin.GoodsTypeAttributeAddController)
	AdminGroups.POST("/goodsTypeAttribute/doAdd", admin.GoodsTypeAttributeDoAddControllerDoAdd)
	AdminGroups.GET("/goodsTypeAttribute/edit", admin.GoodsTypeAttributeEditController)
	AdminGroups.POST("/goodsTypeAttribute/doEdit", admin.GoodsTypeAttributeDoEditController)
	AdminGroups.GET("/goodsTypeAttribute/delete", admin.GoodsTypeAttributeDeleteController)

	AdminGroups.GET("/goods", admin.GoodsController)
	AdminGroups.GET("/goods/delete", admin.DeleteController)
	AdminGroups.GET("/goods/add", admin.GoodsAddController)
	AdminGroups.POST("/goods/doAdd", admin.GoodsDoAddController)
	AdminGroups.GET("/goods/goodsTypeAttribute", admin.GoodsTypeAttrController)
	AdminGroups.GET("/goods/edit", admin.GoodsEditController)
	AdminGroups.POST("/goods/doEdit", admin.GoodsDoEditController)
	AdminGroups.POST("/goods/goodsImageUpload", admin.GoodsImageUpload)
	AdminGroups.POST("/goods/editorImageUpload", admin.EditorImageUpload)
	AdminGroups.GET("/goods/changeGoodsImageColor", admin.ChangeGoodsImageColor)
	AdminGroups.GET("/goods/removeGoodsImage", admin.RemoveGoodsImage)

	AdminGroups.GET("/nav", admin.NavController)
	AdminGroups.GET("/nav/add", admin.NavAddController)
	AdminGroups.POST("/nav/doAdd", admin.NavDoAddController)
	AdminGroups.GET("/nav/edit", admin.NavEditController)
	AdminGroups.POST("/nav/doEdit", admin.NavDoEditController)
	AdminGroups.GET("/nav/delete", admin.NavDeleteController)

	AdminGroups.GET("/setting", admin.SettingController)
	AdminGroups.POST("/setting/doEdit", admin.SettingDoEditController)
}
