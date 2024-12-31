package main

import (
	"nabilwafi/golang_depedency_injection/app"
	"nabilwafi/golang_depedency_injection/controller"
	"nabilwafi/golang_depedency_injection/helper"
	"nabilwafi/golang_depedency_injection/middleware"
	"nabilwafi/golang_depedency_injection/repository"
	"nabilwafi/golang_depedency_injection/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8132",
		Handler: authMiddleware,
	}
}

func main() {
	db := app.NewDB()
	validate := validator.New()
	
	categoryRepositoryImpl := repository.NewCategoryRepository()
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	
	router := app.NewRouter(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	
	server := NewServer(authMiddleware)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}