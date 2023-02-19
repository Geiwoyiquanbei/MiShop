package itying

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckOut(c *gin.Context) {

	cartList := []models.Cart{}

	logic.Cookie.Get(c, "cartList", &cartList)
	var AllPrice float64
	var AllNum int

	OrderList := []models.Cart{}

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			AllPrice += cartList[i].Price * float64(cartList[i].Num)
			OrderList = append(OrderList, cartList[i])
			AllNum += cartList[i].Num
		}
	}

	//2、获取当前用户的收货地址
	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)
	addressList := []models.Address{}
	mysql.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)

	//3、生成签名
	orderSign := logic.Md5(logic.GetRandNum(6))
	session := sessions.Default(c)
	session.Set("orderSign", orderSign)
	session.Save()
	Render(c, "itying/buy/checkout.html", gin.H{
		"orderList":   OrderList,
		"allPrice":    AllPrice,
		"allNum":      AllNum,
		"addressList": addressList,
		"orderSign":   orderSign,
	})
}

func DoCheckOut(c *gin.Context) {
	//0、防止重复提交订单
	orderSignClient := c.PostForm("orderSign")
	session := sessions.Default(c)
	orderSignSession := session.Get("orderSign")
	orderSignServer, ok := orderSignSession.(string)
	if !ok {
		c.Redirect(302, "/")
		return
	}

	if orderSignClient != orderSignServer {
		c.Redirect(302, "/")
		return
	}
	session.Delete("orderSign")
	session.Save()
	// 1、获取用户信息 获取用户的收货地址信息
	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)

	addressResult := []models.Address{}
	mysql.DB.Where("uid = ? AND default_address=1", user.Id).Find(&addressResult)
	if len(addressResult) == 0 {
		c.Redirect(302, "/buy/checkout")
		return
	}
	// 2、获取购买商品的信息
	cartList := []models.Cart{}
	logic.Cookie.Get(c, "cartList", &cartList)
	orderList := []models.Cart{}
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	// 3、把订单信息放在订单表，把商品信息放在商品表
	order := models.Order{
		OrderId:     logic.GetOrderId(),
		Uid:         user.Id,
		AllPrice:    allPrice,
		Phone:       addressResult[0].Phone,
		Name:        addressResult[0].Name,
		Address:     addressResult[0].Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(logic.GetUnix()),
	}

	err := mysql.DB.Create(&order).Error
	//增加数据成功以后可以通过  order.Id
	if err == nil {
		// 把商品信息放在商品对应的订单表
		for i := 0; i < len(orderList); i++ {
			orderItem := models.OrderItem{
				OrderId:      order.Id,
				Uid:          user.Id,
				ProductTitle: orderList[i].Title,
				ProductId:    orderList[i].Id,
				ProductImg:   orderList[i].GoodsImg,
				ProductPrice: orderList[i].Price,
				ProductNum:   orderList[i].Num,
				GoodsVersion: orderList[i].GoodsVersion,
				GoodsColor:   orderList[i].GoodsColor,
			}
			mysql.DB.Create(&orderItem)
		}
	}

	// 4、删除购物车里面的选中数据
	noSelectCartList := []models.Cart{}
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectCartList = append(noSelectCartList, cartList[i])
		}
	}
	logic.Cookie.Set(c, "cartList", noSelectCartList)

	c.Redirect(302, "/buy/pay?orderId="+strconv.Itoa(order.Id))
}

func Pay(c *gin.Context) {

	orderId, err := strconv.Atoi(c.Query("orderId"))
	if err != nil {
		c.Redirect(302, "/")
	}
	//获取用户信息
	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)
	//获取订单信息
	order := models.Order{}
	mysql.DB.Where("id = ?", orderId).Find(&order)
	if order.Uid != user.Id {
		c.Redirect(302, "/")
		return
	}
	//获取订单对应的商品
	orderItems := []models.OrderItem{}
	mysql.DB.Where("order_id = ?", orderId).Find(&orderItems)
	Render(c, "itying/buy/pay.html", gin.H{
		"order":      order,
		"orderItems": orderItems,
	})
}
