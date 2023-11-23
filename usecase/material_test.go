package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/stretchr/testify/assert"
)

func TestMaterialUsecase_Import(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Repo repository.IMaterialRepository
		Ctx  context.Context
	}

	type args struct {
		filepath string
		category string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test importing materials from a csv file",
			fields: fields{
				Repo: repository.NewRepository(db),
				Ctx:  context.Background(),
			},
			args: args{
				filepath: "/home/dera/Documents/Noah CSV Exports/Construction.csv",
				category: "Construction",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mu := MaterialUsecase{
				Repo: tt.fields.Repo,
				Ctx:  tt.fields.Ctx,
			}
			err := mu.BulkImport(tt.args.filepath, tt.args.category)
			assert.Nil(t, err)
		})
	}
}
