package models

import "time"

type User struct {
	ID        int       `json:"id"  gorm:"primary_key:auto_increment"`
	Fullname  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"default:customer"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
