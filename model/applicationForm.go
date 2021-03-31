package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type applicationForm struct {
	gorm.Model

	aFID 				uuid.UUID 	`gorm:"primary_key;unique"`
	name 				string 		`gorm:"unique;not null"`
	brithday			time.Time
	studentId			string		`gorm:"unique;not null"`
	college 			string		`gorm:"unique;not null"`
	profession 			string		`gorm:"unique;not null"`
	phoneNumber			string		`gorm:"unique;not null"`
	qq					string		`gorm:"unique;not null"`
	email 				string 		`gorm:"unique;not null"`
	gender 				string 		`gorm:"unique;not null"`
	selfIntroduction 	string 		`gorm:"unique;not null"`
	blog 				string		
    github 				string
    skills 				string
    elseskills 			string
    expections 			string
	Password			string		`gorm:"not null"`



	CreatedAt   		time.Time
	UpdatedAt   		time.Time
	DeletedAt   		*time.Time

	Role				role		`gorm:"foreignKey:aFId;association_foreignKey:ID"`


	Text string `json:"text"`
}
