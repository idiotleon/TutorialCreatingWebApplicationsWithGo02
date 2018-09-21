package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo/controller"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8080", nil)
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
