package cartdto

type CreateCartRequest struct {
	UserID        int   ` json:"user_id" `
	ProductID     int   ` json:"product_id" `
	Qty           int   ` json:"qty" `
	SubAmount     int   ` json:"sub_amount"`
	TransactionID int64 ` json:"transaction_id"`
}

type CartUpdate struct {
	Qty       int ` json:"qty" `
	SubAmount int ` json:"sub_amount"`
}
