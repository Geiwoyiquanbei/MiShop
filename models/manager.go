package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Role     Role `grom:"foreignKey:RoleId";references:"Id"`
}

func (manager Manager) TableName() string {
	return "manager"
}
