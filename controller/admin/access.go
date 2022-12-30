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

func AccessController(c *gin.Context) {
	accessList := []models.Access{}
	mysql.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}
func AccessAddController(c *gin.Context) {
	//获取顶级模块
	accessList := []models.Access{}
	mysql.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}
func AccessDoAddController(c *gin.Context) {
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err1 := strconv.Atoi(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err2 := strconv.Atoi(c.PostForm("module_id"))
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		logic.ErrorReply(c, "传入参数错误", "/admin/access/add")
		return
	}
	if moduleName == "" {
		logic.ErrorReply(c, "模块名称不能为空", "/admin/access/add")
		return
	}
	access := models.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err5 := mysql.DB.Create(&access).Error
	if err5 != nil {
		logic.ErrorReply(c, "增加数据失败", "/admin/access/add")
		return
	}
	logic.SuccessReply(c, "增加数据成功", "/admin/access")
}
func AccessEditController(c *gin.Context) {
	//获取要修改的数据
	id, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		logic.ErrorReply(c, "参数错误", "/admin/access")
	}
	access := models.Access{Id: id}
	mysql.DB.Find(&access)
	//获取顶级模块
	accessList := []models.Access{}
	mysql.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})
}
func AccessDoEditController(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err2 := strconv.Atoi(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err3 := strconv.Atoi(c.PostForm("module_id"))
	sort, err4 := strconv.Atoi(c.PostForm("sort"))
	status, err5 := strconv.Atoi(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		logic.ErrorReply(c, "传入参数错误", "/admin/access")
		return
	}
	if moduleName == "" {
		logic.ErrorReply(c, "模块名称不能为空", "/admin/access/edit?id="+strconv.Itoa(id))
		return
	}
	access := models.Access{Id: id}
	mysql.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	err := mysql.DB.Save(&access).Error
	if err != nil {
		logic.ErrorReply(c, "修改数据", "/admin/access/edit?id="+strconv.Itoa(id))
	} else {
		logic.SuccessReply(c, "修改数据成功", "/admin/access/edit?id="+strconv.Itoa(id))
	}
}
func AccessDeleteController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logic.ErrorReply(c, "传入数据错误", "/admin/access")
	} else {
		//获取我们要删除的数据
		access := models.Access{Id: id}
		mysql.DB.Find(&access)
		if access.ModuleId == 0 { //顶级模块
			accessList := []models.Access{}
			mysql.DB.Where("module_id = ?", access.Id).Find(&accessList)
			if len(accessList) > 0 {
				logic.ErrorReply(c, "当前模块下面有菜单或者操作，请删除菜单或者操作以后再来删除这个数据", "/admin/access")
			} else {
				mysql.DB.Delete(&access)
				logic.SuccessReply(c, "删除数据成功", "/admin/access")
			}
		} else { //操作 或者菜单
			mysql.DB.Delete(&access)
			logic.SuccessReply(c, "删除数据成功", "/admin/access")
		}
	}
}
