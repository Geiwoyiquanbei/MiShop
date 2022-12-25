package mysql

import "MiShop/models"

func RoleDoAdd(role models.Role) error {
	err := DB.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}
func GetRoleList() []models.Role {
	roleList := []models.Role{}
	DB.Find(&roleList)
	return roleList
}
func RoleDoEdit(role models.Role) error {
	tmp := models.Role{}
	tmp.Id = role.Id
	DB.Find(&tmp)
	tmp.Title = role.Title
	tmp.Description = role.Description
	err := DB.Save(&tmp).Error
	if err != nil {
		return err
	}
	return nil
}
