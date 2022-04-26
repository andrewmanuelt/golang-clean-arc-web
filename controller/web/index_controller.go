package web

import (
	"golang-clean-arc-web/service"
	"net/http"

	"github.com/gorilla/mux"
)

type indexController struct {
	indexService service.IndexService
}

func (controller *indexController) Route(router *mux.Router) {
	router.HandleFunc("/", controller.Index)
}

func (controller *indexController) Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "login", http.StatusFound)
}

func NewIndexController(indexService *service.IndexService) indexController {
	return indexController{
		indexService: *indexService,
	}
}
