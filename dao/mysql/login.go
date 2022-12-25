package mysql

import "MiShop/models"

func Login(username, pass string) []models.Manager {
	UserInfo := []models.Manager{}
	DB.Where("username= ? And password= ?", username, pass).Find(&UserInfo)
	return UserInfo
}
