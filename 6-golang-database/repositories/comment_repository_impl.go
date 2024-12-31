package repositories

import (
	"belajar_golang_db/entities"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository (db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repo *CommentRepositoryImpl) Insert(ctx context.Context, comment entities.Comment) (entities.Comment, error) {
	query := "INSERT INTO Comment(email, comment) VALUES (?,?)"
	res, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repo *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entities.Comment, error) {
	query := "SELECT * FROM Comment WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	comment := entities.Comment{}
	
	if err != nil {
		return comment, err
	}

	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return comment, err
		}

		return comment, nil
	}else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
} 

func (repo *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entities.Comment, error) {
	query := "SELECT * FROM Comment"
	rows, err := repo.DB.QueryContext(ctx, query)
	
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	
	var comments []entities.Comment
	
	for rows.Next() {
		comment := entities.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}