package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home í the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the homepage
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Starting application on port: ", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
