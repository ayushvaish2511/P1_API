package main

import (
	"P1_API/config"
	"P1_API/controller"
	"P1_API/helper"
	"P1_API/repository"
	"P1_API/router"
	"P1_API/service"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Start Server")

	// database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}