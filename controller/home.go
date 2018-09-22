package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "time"

	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo02/model"
	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo02/viewmodel"
)

type home struct {
	homeTemplate         *template.Template
	loginTemplate        *template.Template
	standLocatorTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	// time.Sleep(4 * time.Second)
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		fmt.Println(r.Form)
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if user, err := model.Login(email, password); err == nil {
			log.Printf("User has logged in: %v\n", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("Failed to log user in with email: %v\n, error was: %v\n", email, err)
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
