package usecase

import (
	"context"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/domain"
)

type MaterialUsecase struct {
	Repo repository.IMaterialRepository
	Ctx  context.Context
}
type IMaterialUsecase interface {
	Find()
	Exists()
	Create(name, category, measurement string) domain.Material
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
