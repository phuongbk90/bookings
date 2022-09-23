package main

import (
	"fmt"
	"github.com/phuongbk90/bookings/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting application on port: ", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
