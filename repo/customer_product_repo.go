package repo

import (
	"go_gorm/model"

	"gorm.io/gorm"
)

type CustomerProductRepo interface {
	Create(product *model.Product) error
}

type customerProductRepo struct {
	db *gorm.DB
}

func (c *customerProductRepo) Create(product *model.Product) error {
	// panic("impelement me")

	result := c.db.Create(product).Error
	return result
}

func NewCustomerProductRepository(db *gorm.DB) CustomerProductRepo {
	repo := new(customerProductRepo)
	repo.db = db
	return repo
}
