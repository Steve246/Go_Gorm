package config

import (
	"database/sql"
	"fmt"
	"go_gorm/model"
	"go_gorm/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")         //set DB_HOST=localhost
	dbPort := os.Getenv("DB_PORT")         //set DB_PORT=5432
	dbUser := os.Getenv("DB_USER")         //set DB_USER=postgres
	dbPassword := os.Getenv("DB_PASSWORD") //set DB_PASSWORD=12345678
	dbName := os.Getenv("DB_NAME")         //set DB_NAME=db_enigma_shop_v2

	env := os.Getenv("ENV")

	// urutan url koneksi ke db postgres buat gorm
	// localhost:postgres@12345678:db_enigma_shop_v2/5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	utils.IsError(err)

	if env == "dev" {
		c.Db = db.Debug()
	} else if env == "migration" {
		c.Db = db.Debug()
		err := c.Db.AutoMigrate(&model.Customer{}, &model.UserCredential{}, &model.Address{}, &model.Product{})

		if err != nil {
			return
		}

	} else {
		c.Db = db
	}

}

func (c *Config) DbConn() *gorm.DB {
	return c.Db
}

func (c *Config) DBTutup(*sql.DB) {
	db, err := c.Db.DB()
	utils.IsError(err)
	err = db.Close()
	utils.IsError(err)

}

func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
