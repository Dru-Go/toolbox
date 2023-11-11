package domain

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

type Material struct {
	Id          string `json:"id,omitempty"`
	MaterialId  string `json:"materialId,omitempty"`
	Name        string `json:"name,omitempty"`
	Category    string `json:"category,omitempty"`
	Measurement string `json:"unitOfMeasurement,omitempty"`
}

func CreateUniqueMaterialId(category string) (string, error) {
	b := make([]byte, 3) //equals 8 characters
	rand.Read(b)

	if !Validate(category) {
		return "", errors.New("the provided category does not exist")
	}
	return hex.EncodeToString(b) + category, nil
}
