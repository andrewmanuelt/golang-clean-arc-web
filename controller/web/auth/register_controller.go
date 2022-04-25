package auth

import (
	"fmt"
	"golang-clean-arc-web/config"
	"golang-clean-arc-web/entity"
	"golang-clean-arc-web/helper"
	webModel "golang-clean-arc-web/model/web"
	webService "golang-clean-arc-web/service/web"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type dataRegister struct {
	Username string `validate:"required,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum"`
}

type registerController struct {
	AuthService webService.AuthService
}

func (controller *registerController) Route(router *mux.Router) {
	router.HandleFunc("/register", controller.Register)
}

func (controller *registerController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template := template.Must(template.ParseFiles("./view/auth/register.html"))

		data := &entity.WebResponse{
			Title: config.Env("APP_NAME") + " | Register",
		}

		template.Execute(w, data)
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		data := &dataRegister{
			Username: username,
			Email:    email,
			Password: password,
		}

		valid := helper.Validates(data)

		if valid != nil {
			fmt.Println("validation error : ", valid)

			http.Redirect(w, r, "/register", http.StatusFound)
		}

		controller.AuthService.Register(webModel.RegisterRequest{
			Username: username,
			Email:    email,
			Password: helper.Encrypt([]byte(password)),
		})

		http.Redirect(w, r, "/dashboard", http.StatusFound)
	}
}

func NewRegisterController(authService *webService.AuthService) registerController {
	return registerController{
		AuthService: *authService,
	}
}
