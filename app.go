package main

import (
	"database/sql"
	"go_gorm/config"
	"go_gorm/model"
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

	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}
	}(enigmaDb)

	err := db.AutoMigrate(model.Customer{}) //auto migrate buat bikin table

	if err != nil {
		panic(err)
	}

}
