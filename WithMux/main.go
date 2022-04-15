package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/leksyking/go-server/routes"
)


func main() {
	r := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	routes.ServerRoutes(r)

	http.Handle("/", fileServer)
	http.Handle("/auth", r)
	
	fmt.Println("Server started on port 9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}

