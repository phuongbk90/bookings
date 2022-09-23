package handlers

import (
	"github.com/phuongbk90/bookings/pkg/render"
	"net/http"
)

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the homepage
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
