package main

import (
	"net/http"

	"github.com/IlhamRamadhan-IR/bejalar-restful-api/app"
	"github.com/IlhamRamadhan-IR/bejalar-restful-api/controller"
	"github.com/IlhamRamadhan-IR/bejalar-restful-api/helper"
	"github.com/IlhamRamadhan-IR/bejalar-restful-api/middleware"
	"github.com/IlhamRamadhan-IR/bejalar-restful-api/repository"
	"github.com/IlhamRamadhan-IR/bejalar-restful-api/service"
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
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
