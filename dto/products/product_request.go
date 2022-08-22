package productsdto

type CreateProductRequest struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image int    `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type UpdateProductRequest struct {
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: varchar(255)"`
}
