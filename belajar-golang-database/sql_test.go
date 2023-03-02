package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
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
func TestExecSqlSelect(t *testing.T) {
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
