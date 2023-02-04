package logic

import (
	"MiShop/dao/mysql"
	"MiShop/models"
)

func GetGoodsByCategory(cateId int, goodsType string, limitNum int) []models.Goods {

	//判断cateId 是否是顶级分类
	goodsCate := models.GoodsCate{Id: cateId}
	mysql.DB.Find(&goodsCate)
	var tempSlice []int
	if goodsCate.Pid == 0 { //顶级分类
		//获取顶级分类下面的二级分类
		goodsCateList := []models.GoodsCate{}
		mysql.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)

		for i := 0; i < len(goodsCateList); i++ {
			tempSlice = append(tempSlice, goodsCateList[i].Id)
		}

	}
	tempSlice = append(tempSlice, cateId)

	goodsList := []models.Goods{}
	where := "cate_id in ?"
	switch goodsType {
	case "hot":
		where += " AND is_hot=1"
	case "best":
		where += " AND is_best=1"
	case "new":
		where += " AND is_new=1"
	default:
		break
	}

	mysql.DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Limit(limitNum).Order("sort desc").Find(&goodsList)
	return goodsList
}
