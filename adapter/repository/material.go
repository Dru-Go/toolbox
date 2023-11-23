package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/google/uuid"
)

type MATERIALS struct {
	sq.TableStruct
	ID, MATERIALID, NAME, DESCRIPTION, CATEGORY, UNITOFMEASUREMENT, CREATEDAT, UPDATEDAT sq.StringField
}

type IMaterialRepository interface {
	Find(name string) (domain.Material, error)
	Exists(materialId string) (bool, error)
	Create(name, category, measurement string) (domain.Material, error)
	Import(string, []domain.Material) error
}

func (m MATERIALS) MaterialMapper() func(row *sq.Row) domain.Material {
	return func(row *sq.Row) domain.Material {
		return domain.Material{
			Id:          row.StringField(m.ID),
			Name:        row.StringField(m.NAME),
			MaterialId:  row.StringField(m.MATERIALID),
			Description: row.StringField(m.DESCRIPTION),
			Category:    row.StringField(m.CATEGORY),
			Measurement: row.StringField(m.UNITOFMEASUREMENT),
			CreatedAt:   row.StringField(m.CREATEDAT),
		}
	}
}

func (mt MATERIALS) Values(repo *sql.DB, category string, materials []domain.Material) func(col *sq.Column) {
	materialId, err := domain.CountMaterialWithCategory(repo, category)
	return func(col *sq.Column) {
		for i, m := range materials {
			if err != nil {
				fmt.Printf("Error creating a unique id for material, %v", m)
				return
			}
			col.SetString(mt.ID, uuid.New().String())
			col.SetString(mt.MATERIALID, fmt.Sprintf("%s%06d", domain.Identifier(category), materialId+1+i))
			col.SetString(mt.NAME, m.Name)
			col.SetString(mt.DESCRIPTION, m.Description)
			col.SetString(mt.CATEGORY, category)
			col.SetString(mt.UNITOFMEASUREMENT, m.Measurement)
			col.SetString(mt.CREATEDAT, time.Now().UTC().Format(DateFormat))
			col.SetString(mt.UPDATEDAT, time.Now().Format("0000-00-00 00:00:00"))
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

	count, err := domain.CountMaterialWithCategory(repo.Db, category)
	newMaterialId := fmt.Sprintf("%s%06d", domain.Identifier(category), count+1)
	if err != nil {
		return domain.Material{}, err
	}

	fmt.Printf("Count of materials with category %s is %v \n", category, count)
	result, err := sq.Exec(sq.Log(repo.Db), sq.
		InsertInto(material).
		Columns(material.ID, material.MATERIALID, material.NAME, material.CATEGORY, material.UNITOFMEASUREMENT, material.CREATEDAT, material.UPDATEDAT).
		Values(id, newMaterialId, name, category, measurement, created_at.Format(DateFormat), time.Now().Format("0000-00-00 00:00:00")).SetDialect(sq.DialectMySQL),
	)

	fmt.Println("Results Created", result)
	if err != nil {
		return domain.Material{}, err
	}

	return domain.Material{Id: id, Name: name, MaterialId: fmt.Sprint(count), Category: category, Measurement: measurement}, nil
}

func (repo Repository) Import(category string, materials []domain.Material) error {
	material := sq.New[MATERIALS]("")
	tx, err := repo.Db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	result, err := sq.Exec(tx, sq.
		InsertInto(material).
		ColumnValues(material.Values(repo.Db, category, materials)).SetDialect(sq.DialectMySQL),
	)
	fmt.Println(result.RowsAffected)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
