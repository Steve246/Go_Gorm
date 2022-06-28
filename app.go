package main

import (
	"fmt"
	"go_gorm/config"
	"go_gorm/repo"
	"log"
)

func main() {
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "postgres"
	// dbPassword := "12345678"
	// dbName := "db_enigma_shop_v2"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic(err)
	// }

	// enigmaDb, _ := db.DB()

	// defer func(enigmaDb *sql.DB) {
	// 	err := enigmaDb.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }(enigmaDb)

	// err = enigmaDb.Ping()

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	log.Println("Connected.....")
	// }

	config := config.NewConfigDB()

	db := config.DbConn() //dpt gorm.Db

	enigmaDb, _ := db.DB() //convert ke sql.DB

	// defer func(enigmaDb *sql.DB) {
	// 	err := enigmaDb.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }(enigmaDb)

	//defer atas jadi defer bawah

	defer config.DBTutup(enigmaDb)

	// err := db.AutoMigrate(model.Customer{})
	//auto migrate buat bikin table

	// if err != nil {
	// 	panic(err)
	// }

	//add repo customer_repo

	// repo := repo.NewCustomerRepository(db)

	// customer := model.Customer{
	// 	Id:      "005",
	// 	Name:    "Oban Paor",
	// 	Address: "Bulan",
	// 	Phone:   "0877666333",
	// 	Email:   "Oban.Paor@gmail.com",
	// 	Balance: 100000,
	// }

	// err := repo.Create(&customer)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//nambain update pake struct

	// repo := repo.NewCustomerRepository(db)

	// customer := model.Customer{
	// 	Id: "002",
	// }

	// err := repo.UpdatePakeStruct(&customer, model.Customer{
	// 	Address: "Jakarta",
	// 	Balance: 20000,
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//nambain update pake map

	// repo := repo.NewCustomerRepository(db)

	// customer := model.Customer{
	// 	Id: "002",
	// }

	// err := repo.UpdatePakeMap(&customer, map[string]interface{}{
	// 	"Address":   "",
	// 	"balance":   10000,
	// 	"is_status": false,
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//nambain delete

	// repo := repo.NewCustomerRepository(db)

	// customer := model.Customer{
	// 	Id: "001",
	// }

	// err := repo.Delete(customer.Id) //hard delete

	// err := repo.Delete(&customer) //soft delete

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//ad find by id

	// repo := repo.NewCustomerRepository(db)

	// customer001 := model.Customer{
	// 	Id: "003",
	// }

	// customerFindById, err := repo.FindById(customer001.Id)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println(customerFindById)

	//FindbyAllby

	// repo := repo.NewCustomerRepository(db)
	// customers := []model.Customer{}

	// customers, err := repo.FindAllBy(map[string]interface{}{
	// 	"address": "Bulan",
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	//FindFirstBy

	// repo := repo.NewCustomerRepository(db)

	// customers := model.Customer{}

	// customers, err := repo.FindFirstBy(map[string]interface{}{
	// 	"address": "Jakarta",
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	// fmt.Println("Find First BY:", customers)

	//FindBy

	repo := repo.NewCustomerRepository(db)

	// customers01 := []model.Customer{}

	customers01, err := repo.FindBy("name LIKE ? AND is_status = ?", "%J%", true)

	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Find By:", customers01)

}
