package domain

import (
	"fmt"
	"os"

	"github.com/dru-go/noah-toolbox/utils"
)

type Category struct {
	m map[string]string
}

var CategoryDictionary = Category{m: map[string]string{
	"Construction":   "CN",
	"Electrical":     "EL",
	"Sanitary":       "SN",
	"Mechanical":     "ME",
	"Accessory":      "AC",
	"Paint":          "PA",
	"Surface Finish": "SR",
	"Metal":          "MT",
	"Glass":          "GL",
	"Chemical":       "CH",
	"Stationary":     "ST",
	"Wood":           "WD",
}}

func Identifier(category string) string {
	return CategoryDictionary.m[category]
}

func GetMaterialCategory(materialId string) string {
	key, ok := utils.MapKey(CategoryDictionary.m, materialId[:2])
	if !ok {
		fmt.Fprintf(os.Stdout, "category not found for a material with %s identifier", materialId)
		os.Exit(2)
	}
	return key
}

func Validate(category string) bool {
	_, ok := CategoryDictionary.m[category]
	return ok
}
