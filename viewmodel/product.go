package viewmodel

import (
	"github.com/idiotLeon/TutorialCreatingWebApplicationsWithGo/model"

	"fmt"
)

type ProductViewModel struct {
	Title   string
	Active  string
	Product Product
}

type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageURL         string
	ID               int
}

func productToVM(product model.Product) Product {
	fmt.Println("PricePerLiter of product, product.go: ", product.PricePerLiter)
	return Product{
		Name:             product.Name,
		DescriptionShort: product.DescriptionShort,
		DescriptionLong:  product.DescriptionLong,
		PricePerLiter:    product.PricePerLiter,
		PricePer10Liter:  product.PricePer10Liter,
		Origin:           product.Origin,
		IsOrganic:        product.IsOrganic,
		ImageURL:         product.ImageURL,
		ID:               product.ID,
	}
}
