package admin

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FocusController(c *gin.Context) {
	focusList := []models.Focus{}
	mysql.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}
func FocusAddController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func FocusDoAddController(c *gin.Context) {
	title := c.PostForm("title")
	focusType, err1 := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := strconv.Atoi(c.PostForm("sort"))
	status, err3 := strconv.Atoi(c.PostForm("status"))
	if err1 != nil || err3 != nil {
		logic.ErrorReply(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		logic.ErrorReply(c, "请输入正确的排序值", "/admin/focus/add")
	}
	//上传文件
	focusImgSrc, err4 := logic.UpLoad(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(logic.GetUnix()),
	}
	err5 := mysql.DB.Create(&focus).Error
	if err5 != nil {
		logic.ErrorReply(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		logic.SuccessReply(c, "增加轮播图成功", "/admin/focus")
	}
}
func FocusEditController(c *gin.Context) {
	value := c.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		logic.ErrorReply(c, "增加轮播图失败", "/admin/focus")
	}
	focus := models.Focus{
		Id: id,
	}
	mysql.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}
func FocusDoEditController(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))
	if err1 != nil || err2 != nil || err4 != nil {
		logic.ErrorReply(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		logic.ErrorReply(c, "请输入正确的排序值", "/admin/focus/edit?id="+strconv.Itoa(id))
	}
	//上传文件
	focusImg, _ := logic.UpLoad(c, "focus_img")
	focus := models.Focus{Id: id}
	mysql.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := mysql.DB.Save(&focus).Error
	if err5 != nil {
		logic.ErrorReply(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+strconv.Itoa(id))
	} else {
		logic.SuccessReply(c, "增加轮播图成功", "/admin/focus")
	}
}
func FocusDeleteController(c *gin.Context) {
	id, err :=strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/focus")
	} else {
		focus := models.Focus{Id: id}
		mysql.DB.Delete(&focus)
		//根据自己的需要 要不要删除图片
		// os.Remove("static/upload/20210915/1631694117.jpg")
		logic.SuccessReply(c, "删除数据成功", "/admin/focus")
	}
}
