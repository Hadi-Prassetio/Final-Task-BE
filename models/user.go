package models

import "time"

type User struct {
	ID           int             `json:"id"  gorm:"primary_key:auto_increment"`
	Fullname     string          `json:"fullname" gorm:"type: varchar(255)"`
	Email        string          `json:"email" gorm:"type: varchar(255)"`
	Password     string          `json:"-" gorm:"type: varchar(255)"`
	Status       string          `json:"status" gorm:"default:customer"`
	Profile      ProfileResponse `json:"profile"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transactions []Transaction   `json:"transactions"`
	CreatedAt    time.Time       `json:"-"`
	UpdatedAt    time.Time       `json:"-"`
}

type UserProfile struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

func (UserProfile) TableName() string {
	return "users"
}
