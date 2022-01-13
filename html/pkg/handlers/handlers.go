package handlers

import (
	"net/http"

	"github.com/erincodepirate/go-webapps/html/pkg/config"
	"github.com/erincodepirate/go-webapps/html/pkg/render"
)

// Repo contains repository
var Repo *Repository

// Repository is repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
