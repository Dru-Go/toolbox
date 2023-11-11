package usecase

import "github.com/dru-go/noah-toolbox/domain"

type ProjectUseCase struct {
	Db interface{}
}

type IProjectUseCase interface {
	Find(string) domain.Project
}

// mysql://root:admin@localhost:3306/cookbook
