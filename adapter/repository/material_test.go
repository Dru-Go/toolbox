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
				name:        "Cement #15",
				category:    "Construction",
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
