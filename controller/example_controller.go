package controller

import (
	"encoding/json"
	"golang-clean-arc-web/entity"
	"golang-clean-arc-web/helper"
	"golang-clean-arc-web/model"
	"golang-clean-arc-web/service"
	"net/http"

	"github.com/gorilla/mux"
)

type exampleController struct {
	exampleService service.ExampleService
}

func (controller *exampleController) Route(r *mux.Router) {
	r.HandleFunc("/get", controller.Get)
	r.HandleFunc("/create", controller.Create)
	r.HandleFunc("/update", controller.Update)
	r.HandleFunc("/delete", controller.Delete)
}

type data struct {
	AppName string `validate:"required,alphanum"`
	AppVer  string
}

func (controller *exampleController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	data := controller.exampleService.Get()

	json.NewEncoder(w).Encode(data)
}

func (controller *exampleController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	appName := r.FormValue("app_name")
	appVer := r.FormValue("app_ver")

	data := &data{
		AppName: appName,
		AppVer:  appVer,
	}

	if err_val := helper.Validates(data); err_val != nil {
		res_error := &model.Errors{
			Message: "Validation Error",
			Errors:  err_val,
		}

		json.NewEncoder(w).Encode(res_error)

		return
	}

	last_data := controller.exampleService.Create(appName, appVer)

	res := &model.ExampleResponse{
		AppName: last_data.AppName,
		AppVer:  last_data.AppVer,
	}

	json.NewEncoder(w).Encode(res)
}

func (controller *exampleController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	app_ver := r.FormValue("app_ver")

	update_data := entity.App{
		AppName: "Golang Updates",
		AppVer:  app_ver,
	}

	controller.exampleService.Modify(app_ver, update_data)

	data := controller.exampleService.Get()

	json.NewEncoder(w).Encode(data)
}

func (controller *exampleController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	app_ver := r.FormValue("app_ver")

	controller.exampleService.Erase(app_ver)

	data := controller.exampleService.Get()

	json.NewEncoder(w).Encode(data)
}

func NewExampleController(exampleService *service.ExampleService) exampleController {
	return exampleController{
		exampleService: *exampleService,
	}
}
