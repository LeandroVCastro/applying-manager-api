package company_domain

import (
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/company"
)

type ListCompanies struct {
	CompanyRepository company_repository.CompanyRepositoryInterface
}

func (c ListCompanies) Handle() (companies []*company_repository.SelectNoRelations, errStatus int, err error) {
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
