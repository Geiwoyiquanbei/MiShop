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

func ManagerController(c *gin.Context) {
	managerList := []models.Manager{}
	mysql.DB.Preload("Role").Find(&managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}
func ManagerAddController(c *gin.Context) {
	//获取所有的角色
	roleList := []models.Role{}
	mysql.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}
func ManagerDoAddController(c *gin.Context) {
	roleId, err1 := strconv.Atoi(c.PostForm("role_id"))
	if err1 != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/manager/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	//用户名和密码长度是否合法
	if len(username) < 2 || len(password) < 6 {
		logic.ErrorReply(c, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}
	//判断管理是否存在
	managerList := []models.Manager{}
	mysql.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		logic.ErrorReply(c, "此管理员已存在", "/admin/manager/add")
		return
	}
	//执行增加管理员
	manager := models.Manager{
		Username: username,
		Password: logic.Md5(password),
		Email:    email,
		RoleId:   roleId,
		Mobile:   mobile,
		Status:   1,
	}
	err2 := mysql.DB.Create(&manager).Error
	if err2 != nil {
		logic.ErrorReply(c, "增加管理员失败", "/admin/manager/add")
		return
	}

	logic.SuccessReply(c, "增加管理员成功", "/admin/manager")
}
func ManagerEditController(c *gin.Context) {
	//获取管理员
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	mysql.DB.Find(&manager)
	//获取所有的角色
	roleList := []models.Role{}
	mysql.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}
func ManagerDoEditController(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	if err1 != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err2 := strconv.Atoi(c.PostForm("role_id"))
	if err2 != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	if len(mobile) > 11 {
		logic.ErrorReply(c, "mobile长度不合法", "/admin/manager/edit?id="+strconv.Itoa(id))
		return
	}
	//执行修改
	manager := models.Manager{Id: id}
	mysql.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId

	//注意：判断密码是否为空 为空表示不修改密码 不为空表示修改密码
	if password != "" {
		//判断密码长度是否合法
		if len(password) < 6 {
			logic.ErrorReply(c, "密码的长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+strconv.Itoa(id))
			return
		}
		manager.Password = logic.Md5(password)
	}
	err3 := mysql.DB.Save(&manager).Error
	if err3 != nil {
		logic.ErrorReply(c, "修改数据失败", "/admin/manager/edit?id="+strconv.Itoa(id))
		return
	}
	logic.SuccessReply(c, "修改数据成功", "/admin/manager")
}

func ManagerDeleteController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/manager")
	} else {
		manager := models.Manager{Id: id}
		mysql.DB.Delete(&manager)
		logic.SuccessReply(c, "删除数据成功", "/admin/manager")
	}
}
