package repository

import (
	"database/sql"
	"testing"

	"github.com/dru-go/noah-toolbox/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type args struct {
		name        string
		category    string
		measurement string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Material
		wantErr bool
	}{
		{
			name: "Testing Create statement",
			args: args{
				name:        "Pipe #2",
				category:    "Sanitary",
				measurement: "KG",
			},
		},
		{
			name: "Testing with different category",
			args: args{
				name:        "Cement",
				category:    "Construction",
				measurement: "KG",
			},
		},
		{
			name: "Testing Wrong category",
			args: args{
				name:        "Cement",
				category:    "Constructions",
				measurement: "KG",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			got, err := repo.Create(tt.args.name, tt.args.category, tt.args.measurement)
			assert.NotNil(t, got)
			assert.Nil(t, err)
		})
	}
}

func TestRepository_Find(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				name: "Cement #14",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			got, err := repo.Find(tt.args.name)
			assert.NotNil(t, got)
			assert.Nil(t, err)
		})
	}
}

func TestRepository_Import(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type args struct {
		category  string
		materials []domain.Material
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Bulk Create",
			args: args{
				category: "Construction",
				materials: []domain.Material{
					{
						Name:        "Cement",
						Measurement: "KG",
					},
				},
			},
		},
		{
			name: "Test Bulk With different category",
			args: args{
				category: "Electrical",
				materials: []domain.Material{
					{
						Name:        "Wires",
						Measurement: "KG",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			err := repo.Import(tt.args.category, tt.args.materials)
			assert.Nil(t, err)
		})
	}
}
