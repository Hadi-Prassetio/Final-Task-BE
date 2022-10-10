package transactiondto

type CreateTransaction struct {
	ID     int64 `json:"id"`
	UserID int   `json:"user_id" form:"user_id"`
	Total  int   `gorm:"type: int" json:"total"`
	Qty    int   `gorm:"type: int" json:"qty"`
}

type UpdateTransaction struct {
	Status string `json:"status"`
	UserID int    `json:"user_id"`
	Total  int    `json:"total"`
	Qty    int    `gorm:"type: int" json:"qty"`
}
