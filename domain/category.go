package domain

type Category struct {
	m map[string]string
}

var CategoryDictionary = Category{m: map[string]string{
	"Construction":     "CN",
	"Electrical":       "EL",
	"Sanitary":         "SN",
	"Mechanical":       "ME",
	"Accessories":      "AC",
	"Paints":           "PA",
	"Surface Finishes": "SR",
	"Metals":           "MT",
	"Glasses":          "GL",
	"Chemicals":        "CH",
	"Stationary":       "ST",
}}

func Identifier(category string) string {
	return CategoryDictionary.m[category]
}

func Validate(category string) bool {
	_, ok := CategoryDictionary.m[category]
	return ok
}
