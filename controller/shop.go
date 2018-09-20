package controller

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo/model"
	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo/viewmodel"

	"fmt"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		s.shopTemplate.Execute(w, vm)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	fmt.Println("Length of products: ", len(products))
	vm := viewmodel.NewShopDetail(products)
	s.categoryTemplate.Execute(w, vm)
}
