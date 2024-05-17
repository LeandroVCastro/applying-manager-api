package company_domain

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/company"
)

type ListCompanies struct {
	CompanyRepository company_repository.CompanyRepositoryInterface
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
		CompanyRepository: company_repository.CompanyRepositoryFactory(),
	}
}
