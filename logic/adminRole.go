package logic

import (
	"MiShop/dao/mysql"
	"MiShop/models"
)

func RoleDoAdd(role models.Role) error {
	err := mysql.RoleDoAdd(role)
	if err != nil {
		return err
	}
	return nil
}
func GetRoleList() []models.Role {
	return mysql.GetRoleList()
}
func RoleDoEdit(role models.Role) error {
	err := mysql.RoleDoEdit(role)
	if err != nil {
		return err
	}
	return nil
}
