package repository

import (
	"github.com/bokwoon95/sq"
)

type Transaction struct {
	sq.TableStruct
	ID, MATERIAL_ID, NAME, CATEGORY, MEASUREMENT sq.StringField
}

type ITransactionRepository interface {
}
