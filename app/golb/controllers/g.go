package controllers

import "html/template"

var (
	homeController home
	templates      map[string]*template.Template
)

func init() {
	templates = populateTemplates()
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
