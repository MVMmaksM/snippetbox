package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/MVMmaksM/snippetbox/cmd/web/helpers"
	"github.com/MVMmaksM/snippetbox/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			helpers.NotFound(app, w)
			return
		}

		if r.Method != "GET" {
			w.Header().Set("Allow", http.MethodGet)
			helpers.ClientError(app, w, http.StatusMethodNotAllowed)
			return
		}

		files := []string{
			"../../ui/html/home.page.tmpl",
			"../../ui/html/base.layout.tmpl",
			"../../ui/html/footer.partial.tmpl",
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			helpers.ServerError(app, w, err)
		}
	}
}

func ShowSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Allow", http.MethodGet)
			helpers.ClientError(app, w, http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			helpers.NotFound(app, w)
			return
		}

		fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
	}
}

func CreateSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", http.MethodPost)
			helpers.ClientError(app, w, http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Create snippet"))
	}
}
