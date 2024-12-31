package main

import (
	"nabilwafi/golang_restful_api/app"
	"nabilwafi/golang_restful_api/controller"
	"nabilwafi/golang_restful_api/helper"
	"nabilwafi/golang_restful_api/middleware"
	"nabilwafi/golang_restful_api/repository"
	"nabilwafi/golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:8132",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}