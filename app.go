package main

import (
	"fmt"
	"go_gorm/config"
	"go_gorm/model"
	"go_gorm/repo"
	"log"
)

func main() {

	config := config.NewConfigDB()

	db := config.DbConn() //dpt gorm.Db

	enigmaDb, _ := db.DB() //convert ke sql.DB

	defer config.DBTutup(enigmaDb)

	//sampe sini, pilih dlu opsi ENV, mau create table baru (migration) or dev

	//insert

	// repo := repo.NewCustomerRepository(db)

	// customer01 := model.Customer{
	// 	Id:      "002",
	// 	Name:    "Sukiyaki Goreng",
	// 	Address: "Jakarta",
	// 	Balance: 20000,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "sukisuki",
	// 		Password: "passbeda",
	// 	},
	// 	Email: "suki.enak@gmail.com",
	// 	Phone: "0817895432",
	// }

	// repo.Create(&customer01)

	//delete

	//nambain update pake map

	repo := repo.NewCustomerRepository(db)

	customer02 := model.Customer{
		Id: "001",
	} //cari id

	customer02, err1 := repo.FindById(customer02.Id) //ambil data dari find by id

	if err1 != nil {
		log.Println(err1.Error())
	}

	fmt.Println("Ini Customer02")
	fmt.Println(customer02)

	userCredential01 := model.UserCredential{
		UserName: "bukitbulan",
		Password: "pasbulan",
	} //masukin apa yang mau diupdate

	customer02.UserCredential = userCredential01
	err2 := repo.UpdateBy(&customer02)

	if err2 != nil {
		log.Println(err2.Error())
	}

	//nambain findbyid

}
