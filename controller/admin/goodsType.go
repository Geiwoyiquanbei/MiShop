package admin

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GoodsTypeController(c *gin.Context) {
	goodsTypeList := []models.GoodsType{}
	mysql.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})

}
func GoodsTypeAddController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}
func GoodsTypeDoAddController(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err1 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil {
		logic.ErrorReply(c, "传入的参数不正确", "/admin/goodsType/add")
		return
	}

	if title == "" {
		logic.ErrorReply(c, "标题不能为空", "/admin/goodsType/add")
		return
	}
	goodsType := models.GoodsType{
		Title:       title,
		Description: description,
		Status:      status,
		AddTime:     int(logic.GetUnix()),
	}

	err := mysql.DB.Create(&goodsType).Error
	if err != nil {
		logic.ErrorReply(c, "增加商品类型失败 请重试", "/admin/goodsType/add")
	} else {
		logic.SuccessReply(c, "增加商品类型成功", "/admin/goodsType")
	}
}
func GoodsTypeEditController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		mysql.DB.Find(&goodsType)
		c.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
			"goodsType": goodsType,
		})
	}
}
func GoodsTypeDoEditController(c *gin.Context) {

	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err2 := strconv.Atoi(c.PostForm("status"))
	if err1 != nil || err2 != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/goodsType")
		return
	}

	if title == "" {
		logic.ErrorReply(c, "商品类型的标题不能为空", "/admin/goodsType/edit?id="+strconv.Itoa(id))
	}
	goodsType := models.GoodsType{Id: id}
	mysql.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err3 := mysql.DB.Save(&goodsType).Error
	if err3 != nil {
		logic.ErrorReply(c, "修改数据失败", "/admin/goodsType/edit?id="+strconv.Itoa(id))
	} else {
		logic.SuccessReply(c, "修改数据成功", "/admin/goodsType")
	}

}
func GoodsTypeDeleteController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		mysql.DB.Delete(&goodsType)
		logic.SuccessReply(c, "删除数据成功", "/admin/goodsType")
	}
}
