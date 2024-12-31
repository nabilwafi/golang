package belajar_golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background();

	query := "INSERT INTO Customer(id, name) VALUES ('joko', 'Joko')"
	
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert News Customer");
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql := "SELECT * FROM Customer"
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id);
		fmt.Println("Name :", name)
	}

	defer rows.Close()
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, married, birthdate, created_at FROM Customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthdate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &married ,&birthdate, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Rating:", rating)
		fmt.Println("Balance:", balance)
		fmt.Println("Married:", married)
		if birthdate.Valid {
			fmt.Println("Birthdate:", birthdate.Time)
		}
		fmt.Println("CreatedAt:", createdAt) 
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlQuery := "SELECT username FROM User WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1" 
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login", username)
	}else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlQuery := "SELECT username FROM User Where username = ? AND password = ? LIMIT 1" 
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully Login", username)
	}else {
		fmt.Println("Login Failed!")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "nabil@gmail.com"
	comment := "test komen"

	sqlQuery := "INSERT INTO Comment(email, comment) VALUES (?,?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Insert Id:", insertID)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "INSERT INTO Comment(email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 5; i++ {
		email := "nabil" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komen ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully created comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background();

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	sqlQuery := "INSERT INTO Comment(email, comment) VALUES (?, ?)"
	statement, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 5; i++ {
		email := "nabil" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komen Transaction ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully created comment Id", id)
	}
}