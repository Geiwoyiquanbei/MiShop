package itying

import (
	"MiShop/dao/mysql"
	"MiShop/dao/redis"
	"MiShop/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func Render(c *gin.Context, tpl string, data map[string]interface{}) {

	//1、获取顶部导航
	topNavList := []models.Nav{}
	if hasTopNavList := redis.CacheDb.Get("topNavList", &topNavList); !hasTopNavList {
		mysql.DB.Where("status=1 AND position=1").Find(&topNavList)
		redis.CacheDb.Set("topNavList", topNavList, 60*60)
	}

	//2、获取分类的数据
	goodsCateList := []models.GoodsCate{}

	if hasGoodsCateList := redis.CacheDb.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		//https://gorm.io/zh_CN/docs/preload.html
		mysql.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		redis.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	}

	//3、获取中间导航
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

	renderData := gin.H{
		"topNavList":    topNavList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
	}

	for key, v := range data {
		renderData[key] = v
	}

	c.HTML(http.StatusOK, tpl, renderData)
}
