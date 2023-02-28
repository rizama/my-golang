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
