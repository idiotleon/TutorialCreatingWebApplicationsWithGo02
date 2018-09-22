package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo02/controller"
	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo02/middleware"
	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo02/model"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)
	http.ListenAndServe(":8080", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://idiotleon:idiotLeon@localhost/testdb?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	rootPath, _ := os.Getwd()
	rootTemplatePath := rootPath + "/templates"
	layout := template.Must(template.ParseFiles(rootTemplatePath + "/_layout.html"))
	template.Must(layout.ParseFiles(rootTemplatePath+"/_header.html", rootTemplatePath+"/_footer.html"))
	dir, err := os.Open(rootTemplatePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(rootTemplatePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
