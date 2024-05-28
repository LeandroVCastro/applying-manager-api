package company_repository

import (
	"errors"
	"fmt"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type CompanyRepositoryInterface interface {
	GetById(id uint) *entity.Company
	CreateOrUpdate(id uint, name string, description, website, linkedin, glassdoor, instagram *string) (uint, error)
	ListAll() (companies []*entity.Company, err error)
	Delete(id uint) error
}

type CompanyRepository struct {
	connection *gorm.DB
}

func (repository CompanyRepository) ListAll() (listedCompanies []*entity.Company, err error) {
	result := repository.connection.Select([]string{"id", "name", "description", "website", "linkedin", "glassdoor", "instagram", "created_at", "updated_at"}).Order("id ASC").Find(&listedCompanies)
	if result.Error != nil {
		err = errors.New(result.Error.Error())
		return
	}
	return
}

func (repository CompanyRepository) GetById(id uint) (companyFound *entity.Company) {
	result := repository.connection.Select([]string{"id", "name", "description", "website", "linkedin", "glassdoor", "instagram", "created_at", "updated_at"}).First(&companyFound, "id = ?", id)
	if result.Error != nil {
		return nil
	}
	return
}

func (repository CompanyRepository) Delete(id uint) error {
	var company = entity.Company{}
	result := repository.connection.Where("ID = ?", id).Delete(&company)
	if result.Error != nil {
		err := errors.New(result.Error.Error())
		return err
	}
	return nil
}

func (repository CompanyRepository) CreateOrUpdate(
	id uint,
	name string,
	description *string,
	website *string,
	linkedin *string,
	glassdoor *string,
	instagram *string,
) (savedId uint, err error) {
	companyParams := entity.Company{Name: name}
	if *description != "" {
		companyParams.Description = description
	} else {
		companyParams.Description = nil
	}

	if *website != "" {
		companyParams.Website = website
	} else {
		companyParams.Website = nil
	}

	if *linkedin != "" {
		companyParams.Linkedin = linkedin
	} else {
		companyParams.Linkedin = nil
	}

	if *glassdoor != "" {
		companyParams.Glassdoor = glassdoor
	} else {
		companyParams.Glassdoor = nil
	}

	if *instagram != "" {
		companyParams.Instagram = instagram
	} else {
		companyParams.Instagram = nil
	}

	var result *gorm.DB
	if id != 0 {
		companyParams.ID = id
		result = repository.connection.Model(companyParams).
			Update("name", companyParams.Name).
			Update("description", companyParams.Description).
			Update("website", companyParams.Website).
			Update("linkedin", companyParams.Linkedin).
			Update("glassdoor", companyParams.Glassdoor).
			Update("instagram", companyParams.Instagram)
	} else {
		result = repository.connection.Create(&companyParams)
	}
	if result.Error != nil {
		fmt.Println("Error on createOrUpdate CompanyRepository: " + result.Error.Error())
		err = result.Error
		return
	}
	savedId = companyParams.ID
	return
}

func CompanyRepositoryFactory() CompanyRepository {
	repository := CompanyRepository{
		connection: configs.Connection,
	}
	return repository
}
