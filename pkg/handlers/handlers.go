package handlers

import (
	"net/http"

	"github.com/dmawardi/go-server/pkg/render"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}

