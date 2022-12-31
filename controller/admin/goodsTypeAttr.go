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

func GoodsTypeAttributeController(c *gin.Context) {
	cateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	mysql.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList)
	//获取商品类型属性对应的类型
	goodsType := models.GoodsType{}
	mysql.DB.Where("id=?", cateId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})

}
func GoodsTypeAttributeAddController(c *gin.Context) {
	//获取当前商品类型属性对应的类型id

	cateId, err := strconv.Atoi(c.Query("cate_id"))
	if err != nil {
		logic.ErrorReply(c, "传入的参数不正确", "/admin/goodsType")
		return
	}

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	mysql.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"cateId":        cateId,
	})
}

func GoodsTypeAttributeDoAddControllerDoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err1 := strconv.Atoi(c.PostForm("cate_id"))
	attrType, err2 := strconv.Atoi(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		logic.ErrorReply(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		logic.ErrorReply(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}

	if err3 != nil {
		logic.ErrorReply(c, "排序值不对", "/admin/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(logic.GetUnix()),
	}
	err := mysql.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		logic.ErrorReply(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
	} else {
		logic.SuccessReply(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute?id="+strconv.Itoa(cateId))
	}
}
func GoodsTypeAttributeEditController(c *gin.Context) {

	//获取当前要修改数据的id
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取当前id对应的商品类型属性
	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	mysql.DB.Find(&goodsTypeAttribute)

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	mysql.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeList":      goodsTypeList,
	})
}

func GoodsTypeAttributeDoEditController(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err2 := strconv.Atoi(c.PostForm("cate_id"))
	attrType, err3 := strconv.Atoi(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err4 := strconv.Atoi(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		logic.ErrorReply(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		logic.ErrorReply(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	if err4 != nil {
		logic.ErrorReply(c, "排序值不对", "/admin/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	mysql.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := mysql.DB.Save(&goodsTypeAttr).Error
	if err != nil {
		logic.ErrorReply(c, "修改数据失败", "/admin/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	logic.SuccessReply(c, "需改数据成功", "/admin/goodsTypeAttribute?id="+strconv.Itoa(cateId))
}

func GoodsTypeAttributeDeleteController(c *gin.Context) {
	id, err1 := strconv.Atoi(c.Query("id"))
	cateId, err2 := strconv.Atoi(c.Query("cate_id"))
	if err1 != nil || err2 != nil {
		logic.ErrorReply(c, "传入参数错误", "/admin/goodsType")
	} else {
		goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
		mysql.DB.Delete(&goodsTypeAttr)
		logic.SuccessReply(c, "删除数据成功", "/admin/goodsTypeAttribute?id="+strconv.Itoa(cateId))
	}
}
