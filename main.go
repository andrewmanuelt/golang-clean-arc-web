package main

import (
	"fmt"
	"golang-clean-arc-web/config"
	"golang-clean-arc-web/controller"
	"golang-clean-arc-web/entitiy"
	"golang-clean-arc-web/middleware"
	"golang-clean-arc-web/repository"
	"golang-clean-arc-web/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.Connection()

	db.AutoMigrate(&entitiy.App{})

	exampleRepository := repository.NewExampleRepository(db)

	exampleService := service.NewExampleService(&exampleRepository)

	exampleController := controller.NewExampleController(&exampleService)

	// router := mux.NewRouter()
	sub_router := mux.NewRouter().NewRoute().Subrouter()
	sub_router.Use(middleware.ExampleMiddleware)

	exampleController.Route(sub_router)

	http.ListenAndServe(":9000", sub_router)

	fmt.Println("Server running at 127.0.0.1:9000")
}
