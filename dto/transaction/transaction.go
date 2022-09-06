package transactiondto

type CreateTransaction struct {
	ID     int `json:"id"`
	UserID int `json:"user_id" form:"user_id"`
	Total  int `gorm:"type: int" json:"total"`
}

type UpdateTransaction struct {
	Status string `json:"status"`
	UserID int    `json:"user_id"`
	Total  int    `json:"total"`
}
