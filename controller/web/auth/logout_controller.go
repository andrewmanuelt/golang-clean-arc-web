package auth

import (
	"golang-clean-arc-web/helper"
	webService "golang-clean-arc-web/service/web"
	"net/http"

	"github.com/gorilla/mux"
)

type logoutController struct {
	AuthService webService.AuthService
}

func (controller *logoutController) Route(router *mux.Router) {
	router.HandleFunc("/logout", controller.Logout)
}

func (controller *logoutController) Logout(w http.ResponseWriter, r *http.Request) {
	sess := helper.NewSession()

	sess.DeleteAllSession(w, r)

	http.Redirect(w, r, "/login", http.StatusFound)
}

func NewLogoutController(authService *webService.AuthService) logoutController {
	return logoutController{
		AuthService: *authService,
	}
}
