package gateway

import (
	"github.com/saleh-ghazimoradi/GoGarage/config"
	"net/http"
)

const Version = "0.0.1"

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     config.Appconfig.Env,
		"version": Version,
	}
	if err := jsonResponse(w, http.StatusOK, data); err != nil {
		internalServerError(w, r, err)
	}
}
