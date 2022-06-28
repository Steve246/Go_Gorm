package repo

import (
	"go_gorm/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	UpdatePakeMap(customer *model.Customer, by map[string]interface{}) error
	//pake map

	UpdatePakeStruct(customer *model.Customer, by model.Customer) error //pake struct
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) UpdatePakeMap(customer *model.Customer, by map[string]interface{}) error {

	result := c.db.Model(customer).Updates(by).Error
	return result

}

func (c *customerRepository) UpdatePakeStruct(customer *model.Customer, by model.Customer) error {
	result := c.db.Model(customer).Updates(by)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer).Error

	return result
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
