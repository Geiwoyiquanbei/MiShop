package logic

import "MiShop/models"

//判断购物车里面有没有当前数据
func HasCartData(cartList []models.Cart, currentData models.Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
			return true
		}
	}
	return false
}
