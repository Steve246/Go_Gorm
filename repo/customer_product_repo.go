package repo

import (
	"errors"
	"go_gorm/model"

	"gorm.io/gorm"
)

type CustomerProductRepo interface {
	Create(product *model.Product) error

	FindByIdProduct(id string) (model.Product, error)
	FindAllProduct(by map[string]interface{}) ([]model.Product, error)
	FindByProduct(by string, vals ...interface{}) ([]model.Product, error)

	UpdateByModelProduct(payload *model.Product) error
}

type customerProductRepo struct {
	db *gorm.DB
}

func (c *customerProductRepo) FindByProduct(by string, vals ...interface{}) ([]model.Product, error) {
	var Product []model.Product
	result := c.db.Unscoped().Where(by, vals...).Find(&Product)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return Product, nil
		} else {

			return Product, err
		}

	}
	return Product, nil
}

func (c *customerProductRepo) FindAllProduct(by map[string]interface{}) ([]model.Product, error) {

	var product []model.Product
	result := c.db.Unscoped().Where(by).Find(&product)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}

	return product, nil
}

func (c *customerProductRepo) UpdateByModelProduct(payload *model.Product) error {
	result := c.db.Model(&payload).Updates(payload).Error

	return result
}

func (c *customerProductRepo) FindByIdProduct(id string) (model.Product, error) {

	var product model.Product
	result := c.db.Unscoped().First(&product, "id = ?", id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}

	}

	return product, nil
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
