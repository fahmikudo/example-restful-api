package main

import (
	"github.com/fahmikudo/example-restful-api/application"
	"github.com/fahmikudo/example-restful-api/controller"
	"github.com/fahmikudo/example-restful-api/exception"
	"github.com/fahmikudo/example-restful-api/helper"
	"github.com/fahmikudo/example-restful-api/middleware"
	"github.com/fahmikudo/example-restful-api/repository"
	"github.com/fahmikudo/example-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	db := application.NewDatabase()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepositoryImpl(db)
	categoryService := service.NewCategoryServiceImpl(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
