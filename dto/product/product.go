package productdto

type CreateProductRequest struct {
	Name        string `json:"name" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: varchar(255)" validate:"required"`
	Price       int    `json:"price" form:"price" gorm:"type: varchar(255)" validate:"required"`
	Description string `json:"description" form:"description" gorm:"type: varchar(500)" validate:"required"`
	Image       int    `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: varchar(255)"`
	Description string `json:"description" form:"description" gorm:"type: varchar(500)"`
	Image       int    `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: varchar(255)"`
	Description string `json:"description" form:"description" gorm:"type: varchar(500)"`
	Image       string `json:"image" form:"image" validate:"required"`
}
