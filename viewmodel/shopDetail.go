package viewmodel

import (
	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo02/model"

	"fmt"
)

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	fmt.Println("Length of products, shopDetail: ", len(products))
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(p))
	}
	return result
}
