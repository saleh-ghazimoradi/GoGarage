package gateway

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/GoGarage/config"
	"github.com/saleh-ghazimoradi/GoGarage/docs"
	_ "github.com/saleh-ghazimoradi/GoGarage/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Handlers struct {
}

func Routes(handlers Handlers) http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthCheckHandler)

	swaggerHandler := SetupSwagger()
	router.Handler(http.MethodGet, "/swagger/*any", swaggerHandler)

	return router
}

func SetupSwagger() http.Handler {
	docs.SwaggerInfo.Title = "Golang Web API"
	docs.SwaggerInfo.Description = "This is a web API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", config.Appconfig.ServerAddress)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	return httpSwagger.WrapHandler
}
