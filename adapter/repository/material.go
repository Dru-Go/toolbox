package repository

import (
	"github.com/bokwoon95/sq"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/google/uuid"
)

type MATERIAL struct {
	sq.TableStruct
	ID, MATERIAL_ID, NAME, CATEGORY, MEASUREMENT sq.StringField
}

type IMaterialRepository interface {
	Find(name string) (domain.Material, error)
	Exists(materialId string) (bool, error)
	Create(name, category, measurement string) (domain.Material, error)
}

func (m MATERIAL) Mapper() func(row *sq.Row) domain.Material {
	return func(row *sq.Row) domain.Material {
		return domain.Material{
			Id:          row.StringField(m.ID),
			Name:        row.StringField(m.NAME),
			Category:    row.StringField(m.CATEGORY),
			Measurement: row.StringField(m.MEASUREMENT),
		}
	}
}

func (repo Repository) Exists(materialId string) (bool, error) {
	material := sq.New[MATERIAL]("material")
	query, err := sq.FetchExists(repo.Db, sq.
		SelectOne().
		From(material).
		Where(material.MATERIAL_ID.EqString(materialId)).
		SetDialect(sq.DialectMySQL),
	)
	if err != nil {
		return false, err
	}
	return query, nil
}

func (repo Repository) Find(name string) (domain.Material, error) {
	material := sq.New[MATERIAL]("material")
	query, err := sq.FetchOne(repo.Db, sq.
		From(material).
		Where(material.NAME.EqString(name)).
		SetDialect(sq.DialectMySQL),
		material.Mapper(),
	)
	if err != nil {
		return domain.Material{}, err
	}
	return query, nil
}

// Wrap in a transaction
func (repo Repository) Create(name, category, measurement string) (domain.Material, error) {
	material := sq.New[MATERIAL]("material")
	id := uuid.New().String()

	newMaterialId, err := domain.CreateUniqueMaterialId(category)
	if err != nil {
		return domain.Material{}, err
	}

	_, err = sq.Exec(repo.Db, sq.
		InsertInto(material).
		Columns(material.ID, material.MATERIAL_ID, material.NAME, material.CATEGORY, material.MEASUREMENT).
		Values(id, newMaterialId, name, category, measurement),
	)

	if err != nil {
		return domain.Material{}, err
	}

	return domain.Material{Id: id, Name: name, MaterialId: newMaterialId, Category: category, Measurement: measurement}, nil
}
