package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// Renders template to ResponseWrite using file extension
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Use parse files to grab template files
	parsedTemplate, parseError := template.ParseFiles("./templates/"+ tmpl, "./templates/base.layout.tmpl")
	// If error detected
	if parseError != nil {
		fmt.Println("Error encountered with file parse:",parseError.Error())
		return
	}
	// Execute the template by injecting data
	err := parsedTemplate.Execute(w, nil)

	// If error detected
	if err != nil {
		fmt.Println("Error encountered with render:")
		return
	}
}