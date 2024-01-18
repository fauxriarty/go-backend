package handlers

import (
	"net/http"

	"github.com/fauxriarty/go-backend/pkg/config"
	"github.com/fauxriarty/go-backend/pkg/models"
	"github.com/fauxriarty/go-backend/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	strings := make(map[string]string)
	strings["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: strings,
	})
}
