package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	db := belajar_golang_database.GetConnection()
	commentRepository := NewCommentRepository(db)
	defer db.Close()

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "sam@gmail.com",
		Comment: "Berhasil Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println((result))
}

func TestCommentFindById(t *testing.T) {
	db := belajar_golang_database.GetConnection()
	commentRepository := NewCommentRepository(db)
	defer db.Close()

	ctx := context.Background()
	id := 3

	comment, err := commentRepository.FindById(ctx, int32(id))
	if err != nil {
		panic(err)
	}

	fmt.Println(comment.Id)
	fmt.Println(comment.Email)
	fmt.Println(comment.Comment)
}

func TestCommentFindAll(t *testing.T) {
	db := belajar_golang_database.GetConnection()
	commentRepository := NewCommentRepository(db)
	defer db.Close()

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
