package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "12345678"
	dbName := "db_enigma_shop_v2"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	enigmaDb, _ := db.DB()

	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}

	}(enigmaDb)

	// err = enigmaDb.Ping()

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	log.Println("Connected.....")
	// }

	err = db.AutoMigrate(&Customer{}) //auto migrate buat bikin table
	if err != nil {
		panic(err)
	}

}
