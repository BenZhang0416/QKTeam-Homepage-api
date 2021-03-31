package model

import (
	"api/boot/orm"
)

type role struct {
	ID    int 	  `gorm:"primary_key;unique"`
	Alias string  `gorm:"unique;not null"`
	Name  string  `gorm:"unique;not null"`

	User  []user  `gorm:"foreignKey:RoleId;association_foreignKey:ID"`
}

func (*role) TableName() string {
	return "role"
}

func (*role)Init()  {
	db := 	orm.GetDB()
	db.Create(&role{ID: 0, Alias:"admin", Name:"管理员"})
	db.Create(&role{ID: 1, Alias:"user", Name:"用户"})
}

func (role *role)GetData() map[string]interface{} {
	return map[string]interface{}{
		"roleId": role.ID,
		"alias": role.Alias,
		"name": role.Name,
	}
}