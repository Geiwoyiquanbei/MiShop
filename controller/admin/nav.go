package admin

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func NavController(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	fmt.Println(page)
	//每页显示的数量
	pageSize := 8
	//获取数据
	navList := []models.Nav{}
	mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&navList)

	//获取总数量
	var count int64
	mysql.DB.Table("nav").Count(&count)
	c.HTML(http.StatusOK, "admin/nav/index.html", gin.H{
		"navList": navList,
		//注意float64类型
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
	})
}
func NavAddController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/nav/add.html", gin.H{})
}
func NavDoAddController(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := strconv.Atoi(c.PostForm("position"))
	isOpennew, _ := strconv.Atoi(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if title == "" {
		logic.ErrorReply(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := models.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(logic.GetUnix()),
	}
	err := mysql.DB.Create(&nav).Error
	if err != nil {
		logic.ErrorReply(c, "增加导航失败 请重试", "/admin/nav/add")
	} else {
		logic.SuccessReply(c, "增加导航成功", "/admin/nav")
	}
}
func NavEditController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/nav")
	} else {
		nav := models.Nav{Id: id}
		mysql.DB.Find(&nav)
		fmt.Println(nav)
		c.HTML(http.StatusOK, "admin/nav/edit.html", gin.H{
			"nav": nav,
		})
	}

}
func NavDoEditController(c *gin.Context) {

	id, err1 := strconv.Atoi(c.PostForm("id"))
	if err1 != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/nav")
		return
	}

	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := strconv.Atoi(c.PostForm("position"))
	isOpennew, _ := strconv.Atoi(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if title == "" {
		logic.ErrorReply(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := models.Nav{Id: id}
	mysql.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status
	err2 := mysql.DB.Save(&nav).Error
	if err2 != nil {
		logic.ErrorReply(c, "修改数据失败", "/admin/nav/edit?id="+strconv.Itoa(id))
	} else {
		logic.SuccessReply(c, "修改数据成功", "/admin/nav")
	}

}
func NavDeleteController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/nav")
	} else {
		nav := models.Nav{Id: id}
		mysql.DB.Delete(&nav)
		logic.SuccessReply(c, "删除数据成功", "/admin/nav")
	}
}
