package model

import (
	"api/boot/orm"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type user struct {
	gorm.Model //

	ID 			uuid.UUID 	`gorm:"primary_key;unique"`
	Uesrname 	string 		`gorm:"unique;not null"`
	Email 		string 		`gorm:"unique;not null"`
	Password	string		`gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time

	//外键foreignKey，关联外键association_foreignKey
	Role		role		`gorm:"foreignKey:UserId;association_foreignKey:ID"`

}

func (user *user)TableName(scope *gorm.Scope) string {
	return "user"
}

func (user *user)BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("ID",interface{}(uuid.NewV4()))
	err = scope.SetColumn("CreatedAt", time.Now())
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (user *user)AfterUpdate(scope *gorm.Scope) (err error){
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (user *user) GetData(kind string) map[string]interface{} {
	db := orm.GetDB()
	var role role
	if err := db.Model(user).Related(&role).Find(&role).Error;err != nil{
		panic(err)
	}
	switch kind {
	case "":
		return map[string]interface{} {
			"id": user.ID,

		}
		
	default:
		return make(map[string]interface{})
	}
}
	