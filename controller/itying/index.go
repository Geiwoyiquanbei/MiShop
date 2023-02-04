package itying

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func Index(c *gin.Context) {
	//1、获取顶部导航
	topNavList := []models.Nav{}
	mysql.DB.Where("status=1 AND position=1").Find(&topNavList)
	//2、获取轮播图数据
	focusList := []models.Focus{}
	mysql.DB.Where("status=1 AND focus_type=1").Find(&focusList)
	//3、获取分类的数据
	goodsCateList := []models.GoodsCate{}
	//https://gorm.io/zh_CN/docs/preload.html
	mysql.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
		return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
	}).Find(&goodsCateList)

	//4、获取中间导航
	middleNavList := []models.Nav{}
	mysql.DB.Where("status=1 AND position=2").Find(&middleNavList)

	for i := 0; i < len(middleNavList); i++ {
		relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
		relationIds := strings.Split(relation, ",")
		goodsList := []models.Goods{}
		mysql.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
		middleNavList[i].GoodsItems = goodsList
	}

	//手机

	phoneList := logic.GetGoodsByCategory(1, "best", 8)

	//配件

	otherList := logic.GetGoodsByCategory(9, "all", 1)

	c.HTML(http.StatusOK, "itying/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		"otherList":     otherList,
	})

}
