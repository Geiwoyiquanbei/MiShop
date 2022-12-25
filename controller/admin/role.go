package admin

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func RoleController(c *gin.Context) {
	list := logic.GetRoleList()
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": list,
	})
}
func RoleAddController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func RoleDoAddController(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if len(title) == 0 {
		logic.ErrorReply(c, "角色名不能为空", "/admin/role/add")
		return
	}
	role := models.Role{
		Title:       title,
		Description: description,
		Status:      1,
	}
	err := logic.RoleDoAdd(role)
	if err != nil {
		logic.ErrorReply(c, "增加角色失败，请重试", "/admin/role/add")
		return
	}
	logic.SuccessReply(c, "增加角色成功", "/admin/role")
}
func RoleEditController(c *gin.Context) {
	value := c.Query("id")
	fmt.Println(value)
	id, err := strconv.Atoi(value)
	if err != nil {
		logic.ErrorReply(c, "参数错误", "/admin/role")
		return
	} else {
		role := models.Role{Id: id}
		mysql.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
}
func RoleDoEditController(c *gin.Context) {
	value := c.PostForm("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		logic.ErrorReply(c, "参数错误", "/admin/role/edit")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if len(title) == 0 {
		logic.ErrorReply(c, "角色名不能为空", "/admin/role/edit")
		return
	}
	role := models.Role{
		Id:          id,
		Title:       title,
		Description: description,
	}
	err = logic.RoleDoEdit(role)
	if err != nil {
		logic.ErrorReply(c, "修改角色失败，请重试", "/admin/role/edit?id="+value)
	}
	logic.SuccessReply(c, "修改角色成功", "/admin/role/edit?id="+value)
}
func RoleDeleteController(c *gin.Context) {
	value := c.Query("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		logic.ErrorReply(c, "参数错误", "/admin/role")
		return
	} else {
		role := models.Role{Id: id}
		mysql.DB.Delete(&role)
		logic.SuccessReply(c, "删除角色成功", "/admin/role")
	}
}
