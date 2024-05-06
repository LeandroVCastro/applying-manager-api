package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func StartConnection() {
	fmt.Println("Connecting to database")

	userDatabase := os.Getenv("USER_DATABASE")
	passwordDatabase := os.Getenv("PASSWORD_DATABASE")
	addressDatabase := os.Getenv("ADDRESS_DATABASE")
	portDatabase := os.Getenv("PORT_DATABASE")
	nameDatabase := os.Getenv("NAME_DATABASE")

	dsn := userDatabase + ":" + passwordDatabase + "@tcp(" + addressDatabase + ":" + portDatabase + ")/" + nameDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error on database connection")
	}
	fmt.Println("Database successfully connected")
	Connection = db
}
