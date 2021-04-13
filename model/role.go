package model

import (
	"api/boot/orm"
)

type Role struct {
	ID    int    `gorm:"primary_key;unique"`
	Alias string `gorm:"unique;not null"`
	Name  string `gorm:"unique;not null"`

	User []User `gorm:"foreignKey:RoleId;association_foreignKey:ID"`
}

//初始化
func (*Role) Init() {
	db := orm.GetDB()
	db.Create(&Role{ID: 0, Alias: "admin", Name: "管理员"})
	db.Create(&Role{ID: 1, Alias: "user", Name: "用户"})
}

func (Role *Role) GetData() map[string]interface{} {
	return map[string]interface{}{
		"roleId": Role.ID,
		"alias":  Role.Alias,
		"name":   Role.Name,
	}
}
