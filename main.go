package main

import (
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/helper"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
