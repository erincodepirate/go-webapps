package handlers

import (
	"net/http"

	"github.com/erincodepirate/go-webapps/html/pkg/render"
)

// Home is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
