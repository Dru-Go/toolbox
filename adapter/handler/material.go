package handler

import (
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/dru-go/noah-toolbox/usecase"
)

type MaterialHandler struct {
	Usecase usecase.IMaterialUsecase
}

func NewMaterial(usecase usecase.IMaterialUsecase) MaterialHandler {
	return MaterialHandler{
		Usecase: usecase,
	}
}

func (mh MaterialHandler) Create(name, category, measurement string) (domain.Material, error) {
	// Validate
	// mh.Usecase.Create(name, category, measurement)
	return domain.Material{}, nil
}
