package auth

import (
	"fmt"
	"golang-clean-arc-web/config"
	"golang-clean-arc-web/entity"
	"golang-clean-arc-web/helper"
	webService "golang-clean-arc-web/service/web"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type dataLogin struct {
	Email    string `validate:"email,required"`
	Password string `validate:"alphanum,required"`
}

type loginController struct {
	AuthService webService.AuthService
}

func (controller *loginController) Route(router *mux.Router) {
	router.HandleFunc("/login", controller.Login)
}

func (controller *loginController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template := template.Must(template.ParseFiles("./view/auth/login.html"))

		data := &entity.WebResponse{
			Title: config.Env("APP_NAME") + " | Login",
		}

		template.Execute(w, data)

		return
	}

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		data := &dataLogin{
			Email:    email,
			Password: password,
		}

		helper.Validates(data)

		res := controller.AuthService.Login(email)

		valid := helper.ValidatePassword(res.Password, []byte(password))

		if !valid {
			http.Redirect(w, r, "/login", http.StatusFound)

			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusFound)

		return
	}

	fmt.Println("dashboard")

	http.Redirect(w, r, "/login", http.StatusFound)
}

func NewLoginController(authService *webService.AuthService) loginController {
	return loginController{
		AuthService: *authService,
	}
}
