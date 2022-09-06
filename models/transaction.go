package models

type Transaction struct {
	ID     int         `json:"id" gorm:"primary_key:auto_increment"`
	UserID int         `json:"-"`
	User   UserProfile `json:"user"`
	Carts  []Cart      `json:"carts"`
	Qty    int         `json:"qty"`
	Status string      `json:"status"`
	Total  int         `json:"total"`
}
