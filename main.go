package main

import (
	"firdausalif/go-rest/app"
	"firdausalif/go-rest/controller"
	"firdausalif/go-rest/exception"
	"firdausalif/go-rest/helper"
	"firdausalif/go-rest/repository"
	"firdausalif/go-rest/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	validate := validator.New()
	db := app.NewDB()

	fmt.Println("apllication running")

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
