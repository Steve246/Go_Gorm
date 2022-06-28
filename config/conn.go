package config

import (
	"database/sql"
	"fmt"
	"go_gorm/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// urutan url koneksi ke db postgres buat gorm
	// localhost:postgres@12345678:db_enigma_shop_v2/5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	utils.IsError(err)

	c.Db = db //dbnya gorm.Db

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
