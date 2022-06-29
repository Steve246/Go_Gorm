package main

import (
	"go_gorm/config"
)

func main() {

	config := config.NewConfigDB()

	db := config.DbConn() //dpt gorm.Db

	enigmaDb, _ := db.DB() //convert ke sql.DB

	defer config.DBTutup(enigmaDb)

	// Paging

	// repo := repo.NewCustomerRepository(db)

	// customerPaging, err := repo.Paging(1, 2)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println("result for customerPaging")
	// fmt.Println(customerPaging)

}
