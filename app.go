package main

import (
	"go_gorm/config"
	"go_gorm/model"
	"go_gorm/repo"
)

func main() {

	config := config.NewConfigDB()

	db := config.DbConn() //dpt gorm.Db

	enigmaDb, _ := db.DB() //convert ke sql.DB

	defer config.DBTutup(enigmaDb)

	//sampe sini, pilih dlu opsi ENV, mau create table baru (migration) or dev

	repo := repo.NewCustomerRepository(db)

	customer01 := model.Customer{
		Id:      "002",
		Name:    "Sukiyaki Goreng",
		Address: "Jakarta",
		Balance: 20000,
		UserCredential: model.UserCredential{
			UserName: "sukisuki",
			Password: "passbeda",
		},
		Email: "suki.enak@gmail.com",
		Phone: "0817895432",
	}

	repo.Create(&customer01)

}
