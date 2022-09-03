package models

import "time"

type Profile struct {
	ID        int         `json:"id" gorm:"primary_key:auto_increment"`
	Address   string      `json:"address" gorm:"type: varchar(255)"`
	Phone     string      `json:"phone" gorm:"type: varchar(255)"`
	PostCode  string      `json:"post_code" gorm:"type: varchar(255)"`
	Image     string      `json:"image" gorm:"type: varchar(255)"`
	UserID    int         `json:"user_id"`
	User      UserProfile `json:"user"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}

type ProfileResponse struct {
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	PostCode string `json:"post_code" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	UserID   int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "Profiles"
}
