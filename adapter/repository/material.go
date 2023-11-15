package repository

import (
	"fmt"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/google/uuid"
)

type MATERIALS struct {
	sq.TableStruct
	ID, MATERIALID, NAME, CATEGORY, UNITOFMEASUREMENT, CREATEDAT, UPDATEDAT sq.StringField
}

type IMaterialRepository interface {
	Find(name string) (domain.Material, error)
	Exists(materialId string) (bool, error)
	Create(name, category, measurement string) (domain.Material, error)
}

func (m MATERIALS) MaterialMapper() func(row *sq.Row) domain.Material {
	return func(row *sq.Row) domain.Material {
		return domain.Material{
			Id:          row.StringField(m.ID),
			Name:        row.StringField(m.NAME),
			Category:    row.StringField(m.CATEGORY),
			Measurement: row.StringField(m.UNITOFMEASUREMENT),
			CreatedAt:   row.StringField(m.CREATEDAT),
		}
	}
}

func (repo Repository) Exists(materialId string) (bool, error) {
	material := sq.New[MATERIALS]("material")
	query, err := sq.FetchExists(repo.Db, sq.
		SelectOne().
		From(material).
		Where(material.MATERIALID.EqString(materialId)).
		SetDialect(sq.DialectMySQL),
	)
	if err != nil {
		return false, err
	}
	return query, nil
}

func (repo Repository) Find(name string) (domain.Material, error) {
	material := sq.New[MATERIALS]("material")
	query, err := sq.FetchOne(repo.Db, sq.
		From(material).
		Where(material.NAME.EqString(name)).
		SetDialect(sq.DialectMySQL),
		material.MaterialMapper(),
	)
	if err != nil {
		return domain.Material{}, err
	}
	return query, nil
}

// Wrap in a transaction
func (repo Repository) Create(name, category, measurement string) (domain.Material, error) {
	material := sq.New[MATERIALS]("")
	id := uuid.New().String()
	created_at := time.Now().UTC()
	newMaterialId, err := domain.CreateUniqueMaterialId(category)
	if err != nil {
		return domain.Material{}, err
	}

	result, err := sq.Exec(sq.Log(repo.Db), sq.
		InsertInto(material).
		Columns(material.ID, material.MATERIALID, material.NAME, material.CATEGORY, material.UNITOFMEASUREMENT, material.CREATEDAT, material.UPDATEDAT).
		Values(id, newMaterialId, name, category, measurement, created_at.Format(DateFormat), time.Now().Format("0000-00-00 00:00:00")).SetDialect(sq.DialectMySQL),
	)

	fmt.Println("Results Created", result)
	if err != nil {
		return domain.Material{}, err
	}

	return domain.Material{Id: id, Name: name, MaterialId: newMaterialId, Category: category, Measurement: measurement}, nil
}
