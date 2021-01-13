package router

import (
	"api-customers/application/container"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func Router() *mux.Router {

	customerController := container.CustomerController()

	router := mux.NewRouter()
	router.Use(otelmux.Middleware("jaeger-tracing-go-service"))
	router.HandleFunc("/customers", customerController.GetAll)

	return router
}
