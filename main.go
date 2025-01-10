package main

import (
	"cars/pkg/handlers"
	"fmt"
	"net/http"
)

func main() {
	port := handlers.Args_handler()

	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/car", handlers.HandleCar)
	http.HandleFunc("/filter", handlers.HandleFilter)
	http.HandleFunc("/compare", handlers.HandleCompare)
	http.HandleFunc("/last", handlers.HandleLastViewed)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on port", port[1:])

	http.ListenAndServe(port, nil)
}
