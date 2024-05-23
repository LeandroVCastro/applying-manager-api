package company_domain

import (
	"errors"
	"strconv"

	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/company"
)

type DeleteCompany struct {
	CompanyRepository company_repository.CompanyRepositoryInterface
}

func (c DeleteCompany) Handle(id uint) (errStatus int, err error) {
	if id == 0 {
		err = errors.New("should be provided an ID")
		errStatus = 400
		return
	}
	company := c.CompanyRepository.GetById(id)

	if company == nil {
		err = errors.New("Company not found with ID: " + strconv.Itoa(int(id)))
		errStatus = 404
		return
	}
	errDelete := c.CompanyRepository.Delete(id)
	if errDelete != nil {
		err = errors.New(errDelete.Error())
		errStatus = 500
		return
	}
	return
}

func DeleteCompanyFactory() DeleteCompany {
	return DeleteCompany{
		CompanyRepository: company_repository.CompanyRepositoryFactory(),
	}
}
