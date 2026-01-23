package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/MVMmaksM/snippetbox/config"
)

func main() {
	addr := flag.String("addr", ":9000", "Сетевой адрес")
	flag.Parse()

	app := &config.Application{
		InfoLogger:  log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home(app))
	mux.HandleFunc("/snippet", ShowSnippet(app))
	mux.HandleFunc("/snippet/create", CreateSnippet(app))

	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: app.ErrorLogger,
	}

	app.InfoLogger.Printf("Server started %s:", *addr)
	err := server.ListenAndServe()
	app.ErrorLogger.Fatal(err)

}
