package routers

import (
	"MiShop/controller/front"
	"MiShop/midwares"
	"github.com/gin-gonic/gin"
)

func FrontRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", front.Index)
		defaultRouters.GET("/category:id", front.ProductCategoryController)
		defaultRouters.GET("/detail", front.GoodsDetails)
		defaultRouters.GET("/product/getImgList", front.GetImgList)
		defaultRouters.GET("/cart", front.GetCart)
		defaultRouters.GET("/cart/addCart", front.AddCart)

		defaultRouters.GET("/cart/successTip", front.AddCartSuccess)

		defaultRouters.GET("/cart/decCart", front.DecCart)
		defaultRouters.GET("/cart/incCart", front.IncCart)

		defaultRouters.GET("/cart/changeOneCart", front.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", front.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", front.DelCart)

		defaultRouters.GET("/pass/login", front.LoginController)
		defaultRouters.GET("/pass/captcha", front.GetCaptcha)

		defaultRouters.GET("/pass/registerStep1", front.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", front.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", front.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", front.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", front.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", front.DoRegister)
		defaultRouters.POST("/pass/doLogin", front.DoLogin)
		defaultRouters.GET("/pass/loginOut", front.DoLogOut)

		defaultRouters.GET("/buy/checkout", midwares.UserAuthMidWare, front.CheckOut) //判断用户权限
		defaultRouters.POST("/buy/doCheckout", midwares.UserAuthMidWare, front.DoCheckOut)
		defaultRouters.GET("/buy/pay", midwares.UserAuthMidWare, front.Pay)
		defaultRouters.GET("/buy/orderPayStatus", midwares.UserAuthMidWare, front.OrderPayStatus)

		defaultRouters.GET("/alipay", midwares.UserAuthMidWare, front.AlipayController{}.Alipay)
		defaultRouters.POST("/alipayNotify", front.AlipayController{}.AlipayNotify)
		defaultRouters.GET("/alipayReturn", midwares.UserAuthMidWare, front.AlipayController{}.AlipayReturn)

		defaultRouters.GET("/wxpay", midwares.UserAuthMidWare, front.WxpayController{}.Wxpay)
		defaultRouters.POST("/wxpay/notify", front.WxpayController{}.WxpayNotify)

		defaultRouters.POST("/address/addAddress", midwares.UserAuthMidWare, front.AddAddressController)
		defaultRouters.POST("/address/editAddress", midwares.UserAuthMidWare, front.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", midwares.UserAuthMidWare, front.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", midwares.UserAuthMidWare, front.GetOneAddressList)

		defaultRouters.GET("/user", midwares.UserAuthMidWare, front.UserIndex)
		defaultRouters.GET("/user/order", midwares.UserAuthMidWare, front.UserOrderList)
		defaultRouters.GET("/user/orderinfo", midwares.UserAuthMidWare, front.UserOrderInfo)
	}
}
