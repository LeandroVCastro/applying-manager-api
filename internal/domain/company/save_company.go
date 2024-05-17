package company_domain

import (
	"errors"

	"strconv"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/LeandroVCastro/applying-manager-api/internal/repository"
)

type SaveCompany struct {
	CompanyRepository repository.CompanyRepositoryInterface
}

func (c SaveCompany) Handle(
	id uint,
	name string,
	description *string,
	website *string,
	linkedin *string,
	glassdoor *string,
	instagram *string,
) (savedCompany *entity.Company, errStatus int, err error) {
	if id != 0 {
		if company := c.CompanyRepository.GetById(id); company == nil {
			err = errors.New("company not found by ID '" + strconv.FormatUint(uint64(id), 10) + "'")
			errStatus = 404
			return
		}
		savedId, updatedErr := c.CompanyRepository.CreateOrUpdate(id, name, description, website, linkedin, glassdoor, instagram)
		if updatedErr != nil {
			err = errors.New("error on update company")
			errStatus = 400
			return
		}
		savedCompany = c.CompanyRepository.GetById(savedId)
		return
	}
	savedId, saveErr := c.CompanyRepository.CreateOrUpdate(0, name, description, website, linkedin, glassdoor, instagram)
	if saveErr != nil {
		err = errors.New("error on create company")
		errStatus = 400
		return
	}
	savedCompany = c.CompanyRepository.GetById(savedId)
	return
}

func SaveCompanyFactory() SaveCompany {
	return SaveCompany{
		CompanyRepository: repository.CompanyRepositoryFactory(),
	}
}
