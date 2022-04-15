package routes

import (
	"github.com/gorilla/mux"
	"github.com/leksyking/go-server/controllers"
)

var ServerRoutes = func (route *mux.Router)  {
	route.HandleFunc("/form", controllers.FormHandler).Methods("POST")
	route.HandleFunc("/hello", controllers.HelloHandler).Methods("GET")
	

}

