package model

import "encoding/json"

type Customer struct {
	Id               string `gorm:"primaryKey"`
	Name             string `gorm:"size:50; not null"`
	Address          []Address
	Phone            string
	Email            string `gorm:"unique"`
	Balance          int
	IsStatus         int `gorm:"default:1"` //jangan pake bool, pake int
	UserCredentialID uint
	UserCredential   UserCredential

	BaseModel BaseModel `gorm:"embedded"`
}

func (Customer) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_customers"
}

func (c *Customer) ToString() string {
	customer, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return ""
	}
	return string(customer)
}
