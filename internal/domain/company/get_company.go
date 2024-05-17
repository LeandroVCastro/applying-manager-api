package company_domain

import (
	"errors"
	"strconv"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/LeandroVCastro/applying-manager-api/internal/repository"
)

type GetCompany struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (c GetCompany) Handle(id uint) (company *entity.Company, errStatus int, err error) {
	if id == 0 {
		err = errors.New("should be provided an ID")
		errStatus = 400
		return
	}
	company = c.CompanyRepository.GetById(id)
	if company == nil {
		err = errors.New("Company not found with ID: " + strconv.Itoa(int(id)))
		errStatus = 404
		return
	}
	return

}

func GetCompanyFactory() GetCompany {
	return GetCompany{
		CompanyRepository: repository.CompanyRepositoryFactory(),
	}
}
