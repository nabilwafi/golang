package helper

import "fmt"

var connection string

func init() {
	connection = "MySQL"
	fmt.Println("Init Berjalan")
}

func GetConnection() string {
	return connection
}