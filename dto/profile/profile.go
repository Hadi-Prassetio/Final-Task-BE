package profiledto

import "_waysbeans_/models"

type CreateProfileRequest struct {
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	PostCode string `json:"post_code"`
	UserID   int    `json:"user_id"`
}

type ProfileResponse struct {
	ID       int                `json:"id"`
	Address  string             `json:"address" gorm:"type: varchar(255)"`
	Phone    string             `json:"phone" gorm:"type: varchar(255)"`
	Image    string             `json:"image" gorm:"type: varchar(255)"`
	PostCode string             `json:"post_code"`
	UserID   int                `json:"user_id"`
	User     models.UserProfile `json:"user"`
}
