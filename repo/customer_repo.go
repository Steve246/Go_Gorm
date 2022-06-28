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

//nambain count, group by, and paging

func (c *customerRepository) Count(groupBy string) (int, error) {
	var total int
	result := c.db.Model(&model.Customer{}).Select("Count(*)").Group(groupBy).First(&total)

	if err := result.Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (c *customerRepository) GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&model.Customer{}).Select(selectedBy).Where(whereBy).Group(groupBy).Find(result)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}

	return nil
}

func (c *customerRepository) Paging(page int, itemPerPage int) (interface{}, error) {
	var customers []model.Customer

	offset := itemPerPage * (page - 1)
	res := c.db.Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customers)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return customers, nil
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
