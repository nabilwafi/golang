package belajar_golang_db

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_go_test")

	if err != nil {
		panic(err)
	}
	defer db.Close() // Close Connection
	
	// Use DB
}