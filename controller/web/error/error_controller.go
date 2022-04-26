package error

import (
	"golang-clean-arc-web/config"
	"html/template"
	"net/http"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./view/error/404.html"))

	data := struct {
		Title string
	}{
		config.Env("APP_NAME") + " | 404",
	}

	template.Execute(w, data)
}
