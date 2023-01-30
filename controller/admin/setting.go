package admin

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SettingController(c *gin.Context) {
	setting := models.Setting{}
	mysql.DB.First(&setting)
	c.HTML(http.StatusOK, "admin/setting/index.html", gin.H{
		"setting": setting,
	})
}

func SettingDoEditController(c *gin.Context) {
	setting := models.Setting{Id: 1}
	mysql.DB.Find(&setting)
	if err := c.ShouldBind(&setting); err != nil {
		fmt.Println(err)
		logic.ErrorReply(c, "修改数据失败,请重试", "/admin/setting")
		return
	} else {
		// 上传图片 logo
		siteLogo, err1 := logic.UpLoadImg(c, "site_logo")
		if len(siteLogo) > 0 && err1 == nil {
			setting.SiteLogo = siteLogo
		}
		//上传图片 no_picture
		noPicture, err2 := logic.UpLoadImg(c, "no_picture")
		if len(noPicture) > 0 && err2 == nil {
			setting.NoPicture = noPicture
		}

		err3 := mysql.DB.Save(&setting).Error
		if err3 != nil {
			logic.ErrorReply(c, "修改数据失败", "/admin/setting")
			return
		}
		logic.SuccessReply(c, "修改数据成功", "/admin/setting")
	}
}
