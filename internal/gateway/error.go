package gateway

import (
	"github.com/saleh-ghazimoradi/GoGarage/logger"
	"net/http"
)

func internalServerError(w http.ResponseWriter, r *http.Request) {
	if err := writeJSONError(w, http.StatusInternalServerError, "the server encountered an error"); err != nil {
		logger.Logger.Error("internal server error", "method", r.Method, "path", r.URL.Path, "writeJSONError", err.Error())
	}
}

func badRequestResponse(w http.ResponseWriter, r *http.Request) {
	if err := writeJSONError(w, http.StatusBadRequest, "the request URL is invalid"); err != nil {
		logger.Logger.Warn("bad request", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	}
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	if err := writeJSONError(w, http.StatusNotFound, "the requested resource was not found"); err != nil {
		logger.Logger.Warn("not found", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	}
}

func conflictResponse(w http.ResponseWriter, r *http.Request) {
	if err := writeJSONError(w, http.StatusConflict, "the request URL is conflict"); err != nil {
		logger.Logger.Error("conflict", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	}
}

func unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request) {
	if err := writeJSONError(w, http.StatusUnauthorized, "the request URL is unauthorized"); err != nil {
		logger.Logger.Warn("unauthorized", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	}
}

func unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warn("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	if writeErr := writeJSONError(w, http.StatusUnauthorized, "unauthorized"); writeErr != nil {
		logger.Logger.Error("failed to write unauthorized error response", "method", r.Method, "path", r.URL.Path, "writeError", writeErr.Error())
	}
}

func forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	if err := writeJSONError(w, http.StatusForbidden, "forbidden"); err != nil {
		logger.Logger.Warn("forbidden", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	}
}

func rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	w.Header().Set("Retry_After", retryAfter)
	if err := writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter); err != nil {
		logger.Logger.Warn("rate limit exceeded", "method", r.Method, "path", r.URL.Path)
	}
}
