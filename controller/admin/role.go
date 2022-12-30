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
func RoleAuthController(c *gin.Context) {
	//1、获取角色id
	roleId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/role")
		return
	}
	//2、获取所有的权限
	accessList := []models.Access{}
	mysql.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	roleAccess := []models.RoleAccess{}
	mysql.DB.Where("role_id=?", roleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}
	//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}
	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})
}
func RoleDoAuthController(c *gin.Context) {
	//获取角色id
	roleId, err1 := strconv.Atoi(c.PostForm("role_id"))
	if err1 != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/role")
		return
	}
	//获取权限id  切片
	accessIds := c.PostFormArray("access_node[]")
	//删除当前角色对应的权限
	roleAccess := models.RoleAccess{}
	mysql.DB.Where("role_id=?", roleId).Delete(&roleAccess)
	//增加当前角色对应的权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := strconv.Atoi(v)
		roleAccess.AccessId = accessId
		mysql.DB.Create(&roleAccess)
	}
	fmt.Println("/admin/role/auth?id=?" + strconv.Itoa(roleId))
	// c.String(200, "DoAuth")
	// admin/role/auth?id=9
	logic.SuccessReply(c, "授权成功", "/admin/role/auth?id="+strconv.Itoa(roleId))
}
