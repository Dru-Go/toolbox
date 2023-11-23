package usecase

import (
	"context"
	"fmt"
	"os"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/jedib0t/go-pretty/v6/table"
)

type MaterialUsecase struct {
	Repo repository.IMaterialRepository
	Ctx  context.Context
}
type IMaterialUsecase interface {
	Find()
	Exists()
	Create(name, category, measurement string) domain.Material
	BulkImport(file string) error
	LoadMaterials(file string)
}

func (mu MaterialUsecase) Exists(materialId string) (bool, error) {
	exists, err := mu.Repo.Exists(materialId)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (mu MaterialUsecase) Create(name, category, measurement string) (domain.Material, error) {
	material, err := mu.Repo.Create(name, category, measurement)
	if err != nil {
		return domain.Material{}, err
	}
	return material, nil
}
func (mu MaterialUsecase) BulkImport(filepath string, category string) error {
	csvReader := repository.NewCSVReader(filepath)

	materials, err := repository.ReadCSV[domain.Material](csvReader)
	if err != nil {
		return fmt.Errorf("unable to parse the csv file, %s", err)
	}
	err = mu.Repo.Import(category, materials)
	if err != nil {
		return fmt.Errorf("unable to import the parsed csv file, %s ", err)
	}
	return nil
}

func (mu MaterialUsecase) LoadCSV(filepath string) error {
	csvReader := repository.NewCSVReader(filepath)

	materials, err := repository.ReadCSV[domain.Material](csvReader)
	if err != nil {
		return fmt.Errorf("unable to parse the csv file, %s", err)
	}
	mu.printTable(materials)
	return nil
}

func (m MaterialUsecase) printTable(materials []domain.Material) {
	// Create a new table
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)

	// Define the table header
	tw.AppendHeader(table.Row{
		"MATERIAL_ID",
		"NAME",
		"DESCRIPTION",
		"CATEGORY",
		"MEASUREMENT",
	})

	// Populate the table body with data from the transactions array
	for _, transaction := range materials {
		// Add data to the table body
		tw.AppendRow([]interface{}{
			transaction.MaterialId,
			transaction.Name,
			transaction.Description,
			transaction.Category,
			transaction.Measurement,
		})
	}

	// Set table style
	tw.SetStyle(table.StyleBold)

	// Render the table
	tw.Render()
}
