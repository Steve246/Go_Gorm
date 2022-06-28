package main

import (
	"go_gorm/config"
	"go_gorm/model"
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
	// 	Id:      "004",
	// 	Name:    "Oktabian Elamor",
	// 	Address: "Bulan",
	// 	Phone:   "0877666333",
	// 	Email:   "okta.ela@gmail.com",
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

	repo := repo.NewCustomerRepository(db)

	customer := model.Customer{
		Id: "002",
	}

	err := repo.UpdatePakeMap(&customer, map[string]interface{}{
		"Address":   "",
		"balance":   10000,
		"is_status": false,
	})

	if err != nil {
		log.Println(err.Error())
	}

}
