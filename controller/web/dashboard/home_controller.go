package dashboard

import (
	dashboardService "golang-clean-arc-web/service/dashboard"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type homeController struct {
	HomeService dashboardService.HomeService
}

func (controller *homeController) Route(router *mux.Router) {
	router.HandleFunc("/dashboard", controller.Home)
}

func (controller *homeController) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template := template.Must(template.ParseFiles("./view/dashboard/index.html"))

		data := []struct {
		}{}

		template.Execute(w, data)

		return
	}

	http.Redirect(w, r, "", http.StatusFound)
}

func NewHomeController(homeService *dashboardService.HomeService) homeController {
	return homeController{
		HomeService: *homeService,
	}
}
