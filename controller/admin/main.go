package admin

import (
	"MiShop/dao/mysql"
	"MiShop/models"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func MainController(c *gin.Context) {
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userInfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if ok {
		//1、获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		//2、获取所有的权限
		accessList := []models.Access{}
		mysql.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.order DESC")
		}).Order("sort DESC").Find(&accessList)
		//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
		roleAccess := []models.RoleAccess{}
		mysql.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
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
		fmt.Printf("%#v", accessList)
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/login")
	}

}
func WelcomeController(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
func ChangeController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	// status = ABS(0-1)   1

	// status = ABS(1-1)  0

	err1 := mysql.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}
func ChangeNumController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "传入的参数错误"})
		return
	}
	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")
	err1 := mysql.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	fmt.Println(err1)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "修改数量失败"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "修改数量成功"})
	}
}
