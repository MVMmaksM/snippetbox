package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Server started: http://localhost:9001")
	err := http.ListenAndServe(":9001", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, fmt.Sprintf("Метод %s запрещен для данного роута!", r.Method), http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
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

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, fmt.Sprintf("Метод %s запрещен для данного роута!", r.Method), http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create snippet"))
}
