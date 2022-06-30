package main

import (
	"go_gorm/config"
	"go_gorm/model"
	"go_gorm/repo"

	"go_gorm/utils"
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

	//nambain update pake map --> ubah password, username tapi bikin nambah kolom

	// repo := repo.NewCustomerRepository(db)

	// customer02 := model.Customer{
	// 	Id: "001",
	// } //cari id

	// customer02, err1 := repo.FindById(customer02.Id) //ambil data dari find by id

	// if err1 != nil {
	// 	log.Println(err1.Error())
	// }

	// fmt.Println("Ini Customer02")
	// fmt.Println(customer02)

	// userCredential01 := model.UserCredential{
	// 	UserName: "bukitbulan",
	// 	Password: "pasbulan",
	// } //masukin apa yang mau diupdate

	// customer02.UserCredential = userCredential01
	// err2 := repo.UpdateBy(&customer02)

	// if err2 != nil {
	// 	log.Println(err2.Error())
	// }

	//update with preload --> user_name sama tapi password keubah tanpa nambah baris kolom

	// repo := repo.NewCustomerRepository(db)

	// customer02, err := repo.FindFirstWithPreload(map[string]interface{}{
	// 	"id": "001"},
	// 	"UserCredential",
	// )

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println(customer02)

	// c := customer02.(model.Customer)
	// c.UserCredential.Password = "inirsama"
	// err = repo.UpdateBy(&c)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//insert pake ada address

	// repo := repo.NewCustomerRepository(db)

	// customer01 := model.Customer{
	// 	Id:             "002",
	// 	Name:           "Jandali Zanai",
	// 	Phone:          "0976666",
	// 	Email:          "jandali.zanai@gmail.com",
	// 	Balance:        210000,
	// 	UserCredential: model.UserCredential{UserName: "Jandas", Password: "janzz"},

	// 	Address: []model.Address{
	// 		{
	// 			StreetName: "JL Nin Aja",
	// 			City:       "Jakarta",
	// 			PostalCode: "123",
	// 		},
	// 		{
	// 			StreetName: "JL Braga",
	// 			City:       "Bandung",
	// 			PostalCode: "235",
	// 		},
	// 	},
	// }

	// err := repo.Create(&customer01)
	// utils.IsError(err)

	//tampilin json

	// repo := repo.NewCustomerRepository(db)

	// customer02, err := repo.FindFirstWithPreload(map[string]interface{}{
	// 	"id": "001"},
	// 	"Address",
	// )

	// utils.IsError(err)

	// log.Println(customer02.ToString())

	//Auth Login Check
	// 1. buatlah sebuah usecase authentication login,
	//apabila success balikan informasi customer ==> (authentication_usecase.go)

	// repo := repo.NewCustomerRepository(db)

	// authCheck, err := repo.AuthLogin("fadZong", "pass")

	// utils.IsError(err)

	// log.Println(authCheck)

	// log.Println(authCheck.ToString()) --> error di to string

	//buatlah sebuah usecase member registration dari existing customer yang ada, dengan menambahkan table customer sebuah field isMember (y/n) ==> (member_activation_usecase.go)

	//many to many

	repo := repo.NewCustomerProductRepository(db)

	product01 := model.Product{
		ProductName: "Kacang Manis",
		Customer:    nil,
	}
	err := repo.Create(&product01)
	utils.IsError(err)

	customer01 := model.Customer{
		Id:   "001",
		Name: "Bulan Menerangi",
		Address: []model.Address{
			{
				StreetName: "JL Jalan Aja",
				City:       "Ragunan",
				PostalCode: "12345",
			},
		},
		Phone:   "102030",
		Email:   "bulan.matahari@gmail.com",
		Balance: 10000,
		UserCredential: model.UserCredential{
			UserName: "bulanbulan",
			Password: "hahaihi",
		},
	}

	// customerRepo := repo.NewCustomerProductRepository(db)

	err = repo.AuthLogin(&customer01)
	utils.IsError(err)

}
