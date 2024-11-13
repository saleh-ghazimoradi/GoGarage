package gateway

import (
	"github.com/saleh-ghazimoradi/GoGarage/config"
	"net/http"
)

const Version = "0.0.1"

// healthCheckHandler provides a basic health check for the application
// @Summary Health Check
// @Description Returns the health status, environment, and version of the application, along with failure details if needed
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/healthcheck [get]
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
