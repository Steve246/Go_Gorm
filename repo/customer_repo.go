package repo

import (
	"errors"
	"go_gorm/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	UpdatePakeMap(customer *model.Customer, by map[string]interface{}) error
	//pake map

	UpdatePakeStruct(customer *model.Customer, by model.Customer) error //pake struct

	Delete(customer *model.Customer) error

	FindById(id string) (model.Customer, error)
	//tambain find by id

	FindFirstBy(by map[string]interface{}) (model.Customer, error)
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)
	FindBy(by string, vals ...interface{}) ([]model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) FindFirstBy(by map[string]interface{}) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Unscoped().Where(by).First(&customer)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindBy(by string, vals ...interface{}) ([]model.Customer, error) {
	var customer []model.Customer
	result := c.db.Unscoped().Where(by, vals...).Find(&customer)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return customer, nil
		} else {

			return customer, err
		}

	}
	return customer, nil
}

func (c *customerRepository) FindAllBy(by map[string]interface{}) ([]model.Customer, error) {
	var customer []model.Customer
	result := c.db.Unscoped().Where(by).Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}

	return customer, nil
}

func (c *customerRepository) FindById(id string) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Unscoped().First(&customer, "id = ?", id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}

	}

	return customer, nil
}

// func (c *customerRepository) Delete(id string) error {
// 	result := c.db.Delete(&model.Customer{}, id).Error
// 	return result
// } //kalau mau hard delete

func (c *customerRepository) Delete(customer *model.Customer) error {
	result := c.db.Delete(customer).Error
	return result
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
