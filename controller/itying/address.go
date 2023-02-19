package itying

import (
	"MiShop/dao/mysql"
	"MiShop/logic"
	"MiShop/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddAddressController(c *gin.Context) {
	userinfo := models.User{}
	logic.Cookie.Get(c, "userinfo", &userinfo)
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	var addressNum int64
	addressList := []models.Address{}
	mysql.DB.Table("address").Where("uid = ?", userinfo.Id).Count(&addressNum)
	if addressNum > 10 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "收货地址的数量超过了限制，请编辑以前的收货地址",
		})
		return
	}
	// 3、更新当前用户的所有收货地址的默认收货地址状态为0
	mysql.DB.Table("address").Where("uid = ?", userinfo.Id).Updates(map[string]interface{}{"default_address": 0})

	// 4、增加当前收货地址，让默认收货地址状态是1
	addressResult := models.Address{
		Uid:            userinfo.Id,
		Name:           name,
		Phone:          phone,
		Address:        address,
		DefaultAddress: 1,
	}
	mysql.DB.Create(&addressResult)
	// 5、返回当前用户的所有收货地址返回
	mysql.DB.Table("address").Where("uid=?", userinfo.Id).Order("id desc").Find(&addressList)
	c.JSON(200, gin.H{
		"success":     true,
		"addressList": addressList,
	})
}

// 获取一个收货地址  返回指定收货地址id的收货地址
func GetOneAddressList(c *gin.Context) {
	//1、获取addressId
	addressId, err := strconv.Atoi(c.Query("addressId"))
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误3",
		})
		return
	}
	//2、获取用户id
	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)
	//3、查询当前addressId  userID对应的数据
	addressList := []models.Address{}
	mysql.DB.Where("id = ? AND uid = ?", addressId, user.Id).Find(&addressList)
	if len(addressList) > 0 {
		c.JSON(200, gin.H{
			"success": true,
			"result":  addressList[0],
		})

	} else {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误4",
		})
		return
	}
}

// 编辑收货地址
func EditAddress(c *gin.Context) {
	/*
	   1、获取用户信息以及 表单修改的数据

	   2、更新当前用户的所有收货地址的默认收货地址状态为0

	   3、修改当前收货地址，让默认收货地址状态是1

	 4、查询当前用户的所有收货地址并返回

	*/
	// 1、获取用户信息以及 表单修改的数据
	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)
	id, err := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误2",
		})
		return
	}

	// 2、更新当前用户的所有收货地址的默认收货地址状态为0
	mysql.DB.Table("address").Where("uid = ?", user.Id).Updates(map[string]interface{}{"default_address": 0})

	// 3、修改当前收货地址，让默认收货地址状态是1
	editAddress := models.Address{Id: id}
	mysql.DB.Find(&editAddress)
	editAddress.Name = name
	editAddress.Phone = phone
	editAddress.Address = address
	editAddress.DefaultAddress = 1
	mysql.DB.Save(&editAddress)

	// 4、返回当前用户的所有收货地址返回

	addressList := []models.Address{}
	mysql.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)
	c.JSON(200, gin.H{
		"success": true,
		"result":  addressList,
	})
}

// 修改默认的收货地址
func ChangeDefaultAddress(c *gin.Context) {
	/*
	   1、获取当前用户收货地址id 以及用户id
	   2、更新当前用户的所有收货地址的默认收货地址状态为0
	   3、更新当前收货地址的默认收货地址状态为1
	*/

	user := models.User{}
	logic.Cookie.Get(c, "userinfo", &user)
	addressId, err := strconv.Atoi(c.Query("addressId"))
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误1",
		})
		return
	}
	mysql.DB.Table("address").Where("uid = ?", user.Id).Updates(map[string]interface{}{"default_address": 0})

	mysql.DB.Table("address").Where("uid = ? AND id = ?", user.Id, addressId).Updates(map[string]interface{}{"default_address": 1})

	c.JSON(200, gin.H{
		"success": true,
		"message": "修改数据成功",
	})
}
