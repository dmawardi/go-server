package handlers

import (
	"net/http"

	"github.com/dmawardi/go-server/pkg/config"
	"github.com/dmawardi/go-server/pkg/models"
	"github.com/dmawardi/go-server/pkg/render"
)

// Repository used by handler package
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Create new handler repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Set Repository to parameter
func UpdateRepositoryHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["boo"] = "boodoo"

	render.AltRenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	stringMap["boo"] = "boodoo"

	render.AltRenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
