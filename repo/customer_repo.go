package repo

import (
	"errors"
	"fmt"
	"go_gorm/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	// UpdatePakeMap(customer *model.Customer, by map[string]interface{}) error
	// //pake map

	// UpdatePakeStruct(customer *model.Customer, by model.Customer) error //pake struct

	//update dibawah, bisa keupdate di dua tabel

	UpdateBy(existingCustomer *model.Customer) error //update pake relation

	Delete(customer *model.Customer) error

	// FindById(id string) (model.Customer, error)
	//tambain find by id

	FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error) //findby id dengan relation + preload

	FindFirstBy(by map[string]interface{}) (model.Customer, error)
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)
	FindBy(by string, vals ...interface{}) ([]model.Customer, error)

	//tambain aggregate
	BaseRepositoryAggregation
	BaseRepositoryPaging

	AuthCheck

	//tambain product many2many

	OpenProductForExistingCustomer(customerWithProduct *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

//nambain many2many

func (c *customerRepository) OpenProductForExistingCustomer(customerWithProduct *model.Customer) error {
	result := c.db.Model(&customerWithProduct).Updates(customerWithProduct).Error

	return result
}

//nambain login auth

func (c *customerRepository) AuthLogin(name string, password string) (model.UserCredential, error) {

	var temp model.UserCredential

	if resultName := c.db.Where("user_name = ?", name).First(&temp); resultName.Error != nil {
		return temp, resultName.Error
	}

	if resultPassword := c.db.Where("password = ?", password).First(&temp); resultPassword.Error != nil {
		return temp, resultPassword.Error

	}

	return temp, nil

}

// func (c *customerRepository) AuthCheck(db *gorm.DB, req *model.UserCredential) error {

// 	// user, err := c.db.Model(&model.UserCredential)

// }

//nambain count, group by, and paging

// func (c *customerRepository) Count(groupBy string) (int, error) {
// 	var total int
// 	// result := c.db.Model(&model.Customer{}).Unscoped().Select("Count(*)").Group(groupBy).First(&total)

// 	// result := c.db.Model(&model.Customer{}).Unscoped().Select("count(*)").Find(&total) //bisa return banyak id di dalam

// 	if err := result.Error; err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }

//count diubah jadi dibawah biar ada opsi

// func (c *customerRepository) Count(groupBy string) (int64, error) {
// 	var total int64
// 	var result *gorm.DB

// 	sqlStmt := c.db.Model(&model.Customer{}).Unscoped()

// 	if groupBy == "" {
// 		result = sqlStmt.Count(&total)
// 	} else {
// 		result = sqlStmt.Select("count(*)").Group(groupBy).First(&total)

// 	}

// 	if err := result.Error; err != nil {
// 		return 0, err
// 	}
// 	return total, nil

// }

//diubah lagi jadi bawah

func (c *customerRepository) Count(result interface{}, groupBy string) error {

	var res *gorm.DB

	sqlStmt := c.db.Model(&model.Customer{}).Unscoped()

	if groupBy == "" {
		t, ok := result.(*int64)

		if ok {
			res = sqlStmt.Count(t)
		} else {
			return errors.New("must be int64")
		}

	} else {
		res = sqlStmt.Select(fmt.Sprintf("%s, %s", groupBy, "count(*) as total")).Group(groupBy).Find(result)
	}
	if err := res.Error; err != nil {
		return err
	}

	return nil

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

// func (c *customerRepository) FindById(id string) (model.Customer, error) {
// 	var customer model.Customer
// 	result := c.db.Unscoped().First(&customer, "id = ?", id)

// 	if err := result.Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return customer, nil
// 		} else {
// 			return customer, err
// 		}

// 	}

// 	return customer, nil
// }

func (c *customerRepository) FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error) {
	var customer model.Customer

	result := c.db.Preload(preload).Where(by).First(&customer)

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

// func (c *customerRepository) UpdatePakeMap(customer *model.Customer, by map[string]interface{}) error {

// 	result := c.db.Model(customer).Updates(by).Error
// 	return result

// }

// func (c *customerRepository) UpdatePakeStruct(customer *model.Customer, by model.Customer) error {
// 	result := c.db.Model(customer).Updates(by)

// 	if err := result.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (c *customerRepository) UpdateBy(existingCustomer *model.Customer) error {
	result := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(existingCustomer).Error

	return result

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
