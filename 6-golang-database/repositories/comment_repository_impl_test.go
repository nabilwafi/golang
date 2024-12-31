package repositories

import (
	"belajar_golang_db"
	"belajar_golang_db/entities"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_db.GetConnection())

	ctx := context.Background()
	comment := entities.Comment{
		Email: "repo@gmail.com",
		Comment: "Test Repo",
	}

	res, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_db.GetConnection())

	ctx := context.Background()

	res, err := commentRepository.FindById(ctx, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_db.GetConnection())

	ctx := context.Background()

	res, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}