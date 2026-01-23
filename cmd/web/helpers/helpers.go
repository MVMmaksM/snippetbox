package helpers

import (
	"net/http"

	"github.com/MVMmaksM/snippetbox/config"
)

func ServerError(app *config.Application, w http.ResponseWriter, err error) {
	//trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, err.Error())

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(app *config.Application, w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(app *config.Application, w http.ResponseWriter) {
	ClientError(app, w, http.StatusNotFound)
}
