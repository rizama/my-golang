package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSqlInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "INSERT INTO customer(id, name) VALUES ('1', 'SAM')"

	// ExecContext tidak mengembalikan nilai, jadi cocok untuk proses sql seperti:
	// 1. Insert
	// 2. Update
	// 3. Delete
	_, err := db.ExecContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert new Customer")
}

func TestExecSqlDelete(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "DELETE FROM customer WHERE id='1'"

	// ExecContext tidak mengembalikan nilai, jadi cocok untuk proses sql seperti:
	// 1. Insert
	// 2. Update
	// 3. Delete
	_, err := db.ExecContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Delte Customer")
}

func TestQuerySqlSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name FROM customer"

	// QueryContext mengembalikan nilai, jadi cocok untuk proses sql seperti:
	// 1. Select
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	// Membaca result data
	for rows.Next() {
		var id string
		var name string

		// pembacaan kolom sesuai dengan query select yang digunakan diatas / sesuai urutan pada database
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println(id)
		fmt.Println(name)
	}

	// tutup Rows
	defer rows.Close()

	fmt.Println("Success Select Customer")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	// QueryContext mengembalikan nilai, jadi cocok untuk proses sql seperti:
	// 1. Select
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	// Membaca result data
	for rows.Next() {
		var id string
		var name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		// pembacaan kolom sesuai dengan query select yang digunakan diatas / sesuai urutan pada database
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("========================================")
		fmt.Println("id:", id)
		fmt.Println("name", name)
		fmt.Println("email", email.String)
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("birth_date", birth_date.Time)
		fmt.Println("married", married)
		fmt.Println("created_at", created_at)
	}

	// tutup Rows
	defer rows.Close()

	fmt.Println("Success Select Customer")
}

func TestQuerySqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	querySql := "SELECT username FROM users WHERE username = '" + username + "' AND password ='" + password + "' LIMIT 1"

	fmt.Println(querySql, "Query")
	rows, err := db.QueryContext(ctx, querySql)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		// pembacaan kolom sesuai dengan query select yang digunakan diatas / sesuai urutan pada database
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println(username)
		fmt.Println("Success Login")

	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}

func TestQuerySqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	// menggunakan query parameter variadic
	querySql := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"

	fmt.Println(querySql, "Query")
	rows, err := db.QueryContext(ctx, querySql, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		// pembacaan kolom sesuai dengan query select yang digunakan diatas / sesuai urutan pada database
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println(username)
		fmt.Println("Success Login")

	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}

func TestExecSqlInsertSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "INSERT INTO customer(id, name) VALUES (?, ?)"

	id := "100"
	name := "andre"

	// ExecContext tidak mengembalikan nilai, jadi cocok untuk proses sql seperti:
	// 1. Insert
	// 2. Update
	// 3. Delete
	_, err := db.ExecContext(ctx, querySql, id, name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert new Customer")
}

func TestExecAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	email := "sam@gmail.com"
	comment := "bagus"

	// ExecContext tidak mengembalikan nilai, jadi cocok untuk proses sql seperti:
	// 1. Insert
	// 2. Update
	// 3. Delete
	result, err := db.ExecContext(ctx, querySql, email, comment)
	if err != nil {
		panic(err)
	}

	lastInserId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Id Terakhir :", lastInserId)

	fmt.Println("Success Insert new Comments")
}

// Prepare Statement cocok digunakann ketika akan insert many dalam satu waktu,
// sehingga Exec dan Sql tidak selalu membuat/bertanya ke Pool database
func TestPrepareStatement(t *testing.T) {
	// Get Connection to Pool
	db := GetConnection()
	defer db.Close()

	// Prepare Statement
	ctx := context.Background()
	querySql := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, querySql)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	// execute Prepare Statement
	for i := 0; i < 10; i++ {
		email := "sam" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		// execute with prepare
		// Tidak perlu melampirkan SQL QUery lagi karena sudah builtin ketika create Prepare Statement diatas.
		// Langsung saja masukan ctx, dan parameter
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id:", lastId)

	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Create Transaction Object
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// Do Transaction
	querySql := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "sam" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, querySql, email, comment)
		if err != nil {
			panic(err)
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id:", lastId)

	}

	// after all transaction, do commit
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
