package routers

import (
	"MiShop/controller/itying"
	"github.com/gin-gonic/gin"
)

func FrontRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.Index)
		defaultRouters.GET("/category:id", itying.ProductCategoryController)
		defaultRouters.GET("/detail", itying.GoodsDetails)
		defaultRouters.GET("/product/getImgList", itying.GetImgList)
		defaultRouters.GET("/cart", itying.GetCart)
		defaultRouters.GET("/cart/addCart", itying.AddCart)

		defaultRouters.GET("/cart/successTip", itying.AddCartSuccess)

		defaultRouters.GET("/cart/decCart", itying.DecCart)
		defaultRouters.GET("/cart/incCart", itying.IncCart)

		defaultRouters.GET("/cart/changeOneCart", itying.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", itying.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", itying.DelCart)
	}
}
