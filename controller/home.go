package controller

import (
	"html/template"
	"net/http"

	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}
