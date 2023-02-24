package front

import (
	"MiShop/dao/mysql"
	"MiShop/dao/redis"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Index(c *gin.Context) {
	timeStart := time.Now().UnixNano()
	//1、获取顶部导航

	//2、获取轮播图数据
	focusList := []models.Focus{}
	if hasFocusList := redis.CacheDb.Get("focusList", &focusList); !hasFocusList {
		mysql.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		redis.CacheDb.Set("focusList", focusList, 60*60)
	}
	//3、获取分类的数据
	//4、获取中间导航

	//手机
	phoneList := []models.Goods{}
	if hasPhoneList := redis.CacheDb.Get("phoneList", &phoneList); !hasPhoneList {
		phoneList = logic.GetGoodsByCategory(1, "best", 8)
		redis.CacheDb.Set("phoneList", phoneList, 60*60)
	}

	//配件
	otherList := []models.Goods{}
	if hasOtherList := redis.CacheDb.Get("otherList", &otherList); !hasOtherList {
		otherList = logic.GetGoodsByCategory(9, "all", 1)
		redis.CacheDb.Set("otherList", otherList, 60*60)
	}

	timeEnd := time.Now().UnixNano()

	fmt.Printf("执行时间：%v 毫秒", (timeEnd-timeStart)/1000000)
	Render(c, "front/index/index.html", gin.H{
		"focusList": focusList,
		"phoneList": phoneList,
		"otherList": otherList,
	})

}
