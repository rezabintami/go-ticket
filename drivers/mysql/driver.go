package mysql_driver

import (
	"fmt"
	"log"
	"ticketing/drivers/databases/movies"
	"ticketing/drivers/databases/theater"
	"ticketing/drivers/databases/tickets"
	"ticketing/drivers/databases/topup"
	"ticketing/drivers/databases/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&users.Users{},
		&topup.Topup{},
		&tickets.Tickets{},
		&theater.Theater{},
		&movies.Movie{},
	)

	return db
}
