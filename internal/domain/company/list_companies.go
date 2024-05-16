package company_domain

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/LeandroVCastro/applying-manager-api/internal/repository"
)

type ListCompanies struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (c ListCompanies) Handle() (companies []*entity.Company, errStatus int, err error) {
	companies, err = c.CompanyRepository.ListAll()
	if err != nil {
		errStatus = 400
	}
	return
}

func ListCompaniesFactory() ListCompanies {
	return ListCompanies{
		CompanyRepository: repository.CompanyRepositoryFactory(),
	}
}
