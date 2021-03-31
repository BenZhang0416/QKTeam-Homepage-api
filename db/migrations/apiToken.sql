-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table apiToken
(

    gorm.Model

	userId     uuid.UUID  `gorm:"not null"`
	Token      uuid.UUID  `gorm:"unique;not null"`
	ExpiredAt  datetime

	User       User       `gorm:"foreignKey:UserId;association_foreignKey:ID"`
    primary key (userId)
)