package itying

import (
	"MiShop/dao/mysql"
	"MiShop/dao/redis"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func Index(c *gin.Context) {
	timeStart := time.Now().UnixNano()
	//1、获取顶部导航
	topNavList := []models.Nav{}
	if hasTopNavList := redis.CacheDb.Get("topNavList", &topNavList); !hasTopNavList {
		mysql.DB.Where("status=1 AND position=1").Find(&topNavList)
		redis.CacheDb.Set("topNavList", topNavList, 60*60)
	}
	//2、获取轮播图数据
	focusList := []models.Focus{}
	if hasFocusList := redis.CacheDb.Get("focusList", &focusList); !hasFocusList {
		mysql.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		redis.CacheDb.Set("focusList", focusList, 60*60)
	}
	//3、获取分类的数据
	goodsCateList := []models.GoodsCate{}
	//https://gorm.io/zh_CN/docs/preload.html
	if hasGoodsCateList := redis.CacheDb.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		//https://gorm.io/zh_CN/docs/preload.html
		mysql.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		redis.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	}
	//4、获取中间导航
	middleNavList := []models.Nav{}
	if hasMiddleNavList := redis.CacheDb.Get("middleNavList", &middleNavList); !hasMiddleNavList {
		mysql.DB.Where("status=1 AND position=2").Find(&middleNavList)
		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
			relationIds := strings.Split(relation, ",")
			goodsList := []models.Goods{}
			mysql.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		redis.CacheDb.Set("middleNavList", middleNavList, 60*60)
	}

	for i := 0; i < len(middleNavList); i++ {
		relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
		relationIds := strings.Split(relation, ",")
		goodsList := []models.Goods{}
		mysql.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
		middleNavList[i].GoodsItems = goodsList
	}

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
	c.HTML(http.StatusOK, "itying/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		"otherList":     otherList,
	})

}
