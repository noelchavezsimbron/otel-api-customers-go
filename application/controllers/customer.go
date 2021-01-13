package controllers

import (
	"api-customers/application/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomerController struct {
	service *services.CustomerFinder
}

func NewCustomerController(service *services.CustomerFinder) *CustomerController {
	return &CustomerController{service: service}
}

func (controller CustomerController) GetAll(writer http.ResponseWriter, request *http.Request) {

	customers, err := controller.service.GetAll(request.Context())
	if err != nil {
		responseError(writer, err)
	}

	writer.Header().Add("Status", "200")
	writer.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(writer)
	if err := enc.Encode(customers); err != nil {
		responseError(writer, err)
	}
}

func responseError(writer http.ResponseWriter, err error) {
	writer.Header().Add("Status", "500")
	writer.Write([]byte(fmt.Sprintf(`{"message":"%+v"}`, err)))
}
