package main

import (
	"fmt"
	"golang-clean-arc-web/config"
	"golang-clean-arc-web/controller"
	authController "golang-clean-arc-web/controller/web/auth"
	dashboardController "golang-clean-arc-web/controller/web/dashboard"
	"golang-clean-arc-web/entity"
	"golang-clean-arc-web/repository"
	"golang-clean-arc-web/service"
	dashboardService "golang-clean-arc-web/service/dashboard"
	webService "golang-clean-arc-web/service/web"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.Connection()

	db.AutoMigrate(&entity.App{})
	db.AutoMigrate(&entity.User{})

	exampleRepository := repository.NewExampleRepository(db)
	exampleService := service.NewExampleService(&exampleRepository)
	exampleController := controller.NewExampleController(&exampleService)

	dashboardRepository := repository.NewDashboardRepository(db)
	homeService := dashboardService.NewHomeService(&dashboardRepository)
	homeController := dashboardController.NewHomeController(&homeService)

	authRepository := repository.NewAuthRepository(db)
	authService := webService.NewAuthService(&authRepository)

	loginController := authController.NewLoginController(&authService)
	registerController := authController.NewRegisterController(&authService)

	router := mux.NewRouter()

	static := http.FileServer(http.Dir("assets/"))

	router.PathPrefix("/asset").Handler(http.StripPrefix("/asset", static))

	exampleController.Route(router)
	loginController.Route(router)
	registerController.Route(router)
	homeController.Route(router)

	fmt.Println("Server running at 127.0.0.1:9000")

	http.ListenAndServe(":9000", router)
}
