package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/MVMmaksM/snippetbox/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if r.Method != "GET" {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, fmt.Sprintf("Метод %s запрещен для данного роута!", r.Method), http.StatusMethodNotAllowed)
			return
		}

		files := []string{
			"../../ui/html/home.page.tmpl",
			"../../ui/html/base.layout.tmpl",
			"../../ui/html/footer.partial.tmpl",
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			app.ErrorLogger.Println(err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			app.ErrorLogger.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func ShowSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, fmt.Sprintf("Метод %s запрещен для данного роута!", r.Method), http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
	}
}

func CreateSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, fmt.Sprintf("Метод %s запрещен для данного роута!", r.Method), http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Create snippet"))
	}
}
