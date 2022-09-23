package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/"+"base.layout.tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
}
