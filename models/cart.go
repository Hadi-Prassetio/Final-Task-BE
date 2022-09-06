package models

type Cart struct {
	ID            int                `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int                `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product       ProductTransaction `json:"product"`
	TransactionID int                `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty           int                `json:"qty" form:"qty" gorm:"type:int"`
	SubAmount     int                `json:"sub_amount" form:"sub_amount" gorm:"type:int"`
}
