package model

import ()

type View struct {
	SeeAll bool `gorm:"unique;not null"`
}

func (view *View) TableName() string {
	return "view"
}
