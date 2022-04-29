package main

import (
	"fmt"
	"golang-clean-arc-web/config"
	"golang-clean-arc-web/controller"
	webController "golang-clean-arc-web/controller/web"
	authController "golang-clean-arc-web/controller/web/auth"
	dashboardController "golang-clean-arc-web/controller/web/dashboard"
	errorController "golang-clean-arc-web/controller/web/error"
	"golang-clean-arc-web/entity"
	"golang-clean-arc-web/middleware"
	"golang-clean-arc-web/repository"
	"golang-clean-arc-web/service"
	dashboardService "golang-clean-arc-web/service/dashboard"
	webService "golang-clean-arc-web/service/web"
	"net/http"

	"github.com/gorilla/mux"
)

type customFilesys struct {
	fs http.FileSystem
}

func (cfs customFilesys) Open(name string) (result http.File, err error) {
	f, err := cfs.fs.Open(name)

	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()

	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		f.Close()

		errfile, err := cfs.fs.Open("/index.html")

		if err != nil {
			return nil, err
		}

		return errfile, nil
	}

	return f, nil
}

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

	indexRepository := repository.NewIndexRepository(db)
	indexService := service.NewIndexService(&indexRepository)
	indexController := webController.NewIndexController(&indexService)

	loginController := authController.NewLoginController(&authService)
	registerController := authController.NewRegisterController(&authService)
	logoutController := authController.NewLogoutController(&authService)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(errorController.Error404)

	static := http.FileServer(customFilesys{http.Dir("assets/")})

	router.PathPrefix("/asset").Handler(http.StripPrefix("/asset", static))

	sub_router := router.NewRoute().Subrouter()
	sub_router.Use(middleware.AuthMiddleware)

	exampleController.Route(router)
	loginController.Route(router)
	registerController.Route(router)
	homeController.Route(sub_router)
	logoutController.Route(sub_router)
	indexController.Route(router)

	fmt.Println("Server running at 127.0.0.1:9000")

	http.ListenAndServe(":9000", router)
}
