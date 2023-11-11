package usecase

import "github.com/dru-go/noah-toolbox/domain"

type CompanyUseCase interface {
	Find(string) domain.Company
}
