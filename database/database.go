package database

import "fmt"

var connection string

func init() {
	fmt.Println("function init dijalankan otomatis saat package database di import")
	connection = "Mysql"
}

func GetDatabase() string {
	return connection
}
