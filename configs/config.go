package configs

import (
	"log"
	"ticketing/models/movies"
	"ticketing/models/theater"
	"ticketing/models/tickets"
	"ticketing/models/topup"
	"ticketing/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	MysqlDsn = `root:@/ticketing?parseTime=True&charset=utf8`
	// MysqlDsn = `root:@/ticketing@tcp(127.0.0.1:3306)/amartha?charset=utf8mb4&parseTime=True&loc=Local`
)

func InitDB() {
	dsn := MysqlDsn
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(`failed to connect to mysql: ` + err.Error())
	}
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&users.Users{}, &topup.TopUp{}, &tickets.Tickets{}, &theater.Theater{}, &movies.Movie{})
}
