package domain

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
)

type Material struct {
	Id          string `json:"id,omitempty"`
	MaterialId  string `json:"materialId,omitempty"`
	Name        string `json:"name,omitempty"`
	Category    string `json:"category,omitempty"`
	Measurement string `json:"unitOfMeasurement,omitempty"`
	CreatedAt   string `json:"CreatedAt,omitempty"`
}

func randInt(min, max int) []byte {
	unique := min + rand.Intn(max-min)
	return []byte(fmt.Sprint(unique))
}
func CreateUniqueMaterialId(category string) (string, error) {
	if !Validate(category) {
		return "", errors.New("the provided category does not exist")
	}
	return Identifier(category) + hex.EncodeToString(randInt(100, 999)), nil
}
