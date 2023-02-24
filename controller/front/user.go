package front

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

func UserIndex(c *gin.Context) {
	var tpl string = "front/user/welcome.html"
	Render(c, tpl, gin.H{})
}

func UserOrderList(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//pageSize

	pageSize := 2
	//获取用户信息
	userInfo := models.User{}
	logic.Cookie.Get(c, "userinfo", &userInfo)

	orderList := []models.Order{}

	//模糊查询
	where := "uid =" + strconv.Itoa(userInfo.Id)
	keywords := c.Query("keywords")
	if keywords != "" {
		//查询
		orderItemList := []models.OrderItem{}
		mysql.DB.Where("product_title like ?", "%"+keywords+"%").Find(&orderItemList)
		var str string
		// 字符串：   12,12,22
		for i := 0; i < len(orderItemList); i++ {
			if i == 0 {
				str += strconv.Itoa(orderItemList[i].OrderId)
			} else {
				str += "," + strconv.Itoa(orderItemList[i].OrderId)
			}
		}
		where += " AND id in (" + str + ")"
	}
	//按照状态筛选订单
	orderStatus, statusErr := strconv.Atoi(c.Query("orderStatus"))
	if statusErr == nil && orderStatus >= 0 {
		where += " AND order_status=" + strconv.Itoa(orderStatus)
	} else {
		orderStatus = -1
	}
	mysql.DB.Where(where).Preload("order_item").Order("add_time desc").Find(&orderList)

	var count int64
	mysql.DB.Table("order").Where(where).Find(&count)

	var tpl string = "front/user/order.html"
	Render(c, tpl, gin.H{
		"order":      orderList,
		"page":       pageSize,
		"keywords":   keywords,
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
	})
}

func UserOrderInfo(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.Redirect(302, "/user/order")
	}
	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)
	order := []models.Order{}
	mysql.DB.Where("id=? And uid=?", id, user.Id).Preload("OrderItem").Find(&order)

	if len(order) == 0 {
		c.Redirect(302, "/user/order")
		return
	}
	var tpl = "front/user/order_info.html"
	Render(c, tpl, gin.H{
		"order": order[0],
	})
}
