package routes

import (
	"github.com/gorilla/mux"
	"goRest/controllers"
	"log"
	"net/http"
)

func RegisterEmployeeRoutes(router *mux.Router) {
	router.Handle("/employee/", appHandler(controllers.GetEmployees)).Methods("GET")
	router.Handle("/employee/{id}/", appHandler(controllers.GetEmployeeById)).Methods("GET")
	router.Handle("/employee/city/{city}/", appHandler(controllers.GetEmployeesByCity)).Methods("GET")
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
