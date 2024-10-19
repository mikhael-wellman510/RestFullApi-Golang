package main

import (
	"fmt"
	"net/http"
	app "restfull-api/App"
	repository "restfull-api/Repository"
	service "restfull-api/Service"
	"restfull-api/controller"
	"restfull-api/helper"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	validate := validator.New()
	db := app.NewDb()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting server:", err)
		helper.PanicIfErr(err)
	} else {
		fmt.Println("Running ....")
	}

}
