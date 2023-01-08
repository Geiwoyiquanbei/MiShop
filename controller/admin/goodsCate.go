package admin

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GoodsCateController(c *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	mysql.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}
func GoodsCateAddController(c *gin.Context) {
	catesList := []models.GoodsCate{}
	mysql.DB.Where("pid=?", 0).Find(&catesList)
	c.HTML(200, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": catesList,
	})
}
func GoodsCateDoAddController(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := strconv.Atoi(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := strconv.Atoi(c.PostForm("sort"))
	status, err3 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		logic.ErrorReply(c, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err2 != nil {
		logic.ErrorReply(c, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := logic.UpLoadImg(c, "cate_img")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(logic.GetUnix()),
	}
	err := mysql.DB.Create(&goodsCate).Error
	if err != nil {
		logic.ErrorReply(c, "增加数据失败", "/admin/goodsCate/add")
		return
	}
	logic.SuccessReply(c, "增加数据成功", "/admin/goodsCate")
}
func GoodsCateEditController(c *gin.Context) {
	//获取要修改的数据
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入参数错误", "/admin/goodsCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	mysql.DB.Find(&goodsCate)

	//获取顶级分类
	goodsCateList := []models.GoodsCate{}
	mysql.DB.Where("pid = 0").Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})

}
func GoodsCateDoEditController(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	pid, err2 := strconv.Atoi(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		logic.ErrorReply(c, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err3 != nil {
		logic.ErrorReply(c, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := logic.UpLoadImg(c, "cate_img")
	goodsCate := models.GoodsCate{Id: id}
	mysql.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}
	err := mysql.DB.Save(&goodsCate).Error
	if err != nil {
		logic.ErrorReply(c, "修改失败", "/admin/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	logic.SuccessReply(c, "修改成功", "/admin/goodsCate")
}
func GoodsCateDeleteController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/goodsCate")
	} else {
		//获取我们要删除的数据
		goodsCate := models.GoodsCate{Id: id}
		mysql.DB.Find(&goodsCate)
		if goodsCate.Pid == 0 { //顶级分类
			goodsCateList := []models.GoodsCate{}
			mysql.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)
			if len(goodsCateList) > 0 {
				logic.ErrorReply(c, "当前分类下面子分类，请删除子分类作以后再来删除这个数据", "/admin/goodsCate")
			} else {
				mysql.DB.Delete(&goodsCate)
				logic.SuccessReply(c, "删除数据成功", "/admin/goodsCate")
			}
		} else { //操作 或者菜单
			mysql.DB.Delete(&goodsCate)
			logic.SuccessReply(c, "删除数据成功", "/admin/goodsCate")
		}

	}
}
