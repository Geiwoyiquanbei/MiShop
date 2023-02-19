package routers

import (
	"MiShop/controller/itying"
	"MiShop/midwares"
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

		defaultRouters.GET("/pass/login", itying.LoginController)
		defaultRouters.GET("/pass/captcha", itying.GetCaptcha)

		defaultRouters.GET("/pass/registerStep1", itying.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", itying.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", itying.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", itying.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", itying.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", itying.DoRegister)
		defaultRouters.POST("/pass/doLogin", itying.DoLogin)
		defaultRouters.GET("/pass/loginOut", itying.DoLogOut)

		defaultRouters.GET("/buy/checkout", midwares.UserAuthMidWare, itying.CheckOut) //判断用户权限
		defaultRouters.POST("/buy/doCheckout", midwares.UserAuthMidWare, itying.DoCheckOut)
		defaultRouters.GET("/buy/pay", midwares.UserAuthMidWare, itying.Pay)

		defaultRouters.POST("/address/addAddress", midwares.UserAuthMidWare, itying.AddAddressController)
		defaultRouters.POST("/address/editAddress", midwares.UserAuthMidWare, itying.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", midwares.UserAuthMidWare, itying.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", midwares.UserAuthMidWare, itying.GetOneAddressList)
	}
}
