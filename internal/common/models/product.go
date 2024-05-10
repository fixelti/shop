package models

import "time"

type ProductEntity struct {
	ID          uint       `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Price       uint       `db:"price"`
	ImageURL    string     `db:"image_url"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdateAt    time.Time  `db:"updated_at"`
	DeleteAt    *time.Time `db:"deleted_at"`
}

type ProductDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	ImageURL    string `json:"image_url"`
}

func (productEntity ProductEntity) ToDTO() ProductDTO {
	return ProductDTO{
		ID:          productEntity.ID,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
		ImageURL:    productEntity.ImageURL,
	}
}
