package company_domain

import (
	"errors"
	"testing"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/company"
	"github.com/stretchr/testify/assert"
)

var description string = "teste"
var website string = "testewebsite"
var linkedin string = "testelinkedin"
var glassdoor string = "testeglassdoor"
var instagram string = "testeinstagram"
var expectedCompany = &entity.Company{
	ID:          1,
	Name:        "Company test name",
	Description: &description,
	Website:     &website,
	Linkedin:    &linkedin,
	Glassdoor:   &glassdoor,
	Instagram:   &instagram,
}

func TestSaveCompanyDomain(t *testing.T) {
	t.Run("Should return error 404 when an ID is provided and company is not found", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		var expectedCompany *entity.Company
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		saveCompanyDomain := company_domain.SaveCompany{CompanyRepository: mockCompanyRepository}
		createdCompany, errStatus, err := saveCompanyDomain.Handle(1, "name teste", &description, &website, &linkedin, &glassdoor, &instagram)
		assert.Nil(t, createdCompany)
		assert.Equal(t, errStatus, 404)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return error when update fails", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		mockCompanyRepository.On(
			"CreateOrUpdate",
			uint(1),
			"Company test name",
			&description,
			&website,
			&linkedin,
			&glassdoor,
			&instagram,
		).Return(0, errors.New("error on update"))
		saveCompanyDomain := company_domain.SaveCompany{CompanyRepository: mockCompanyRepository}
		createdCompany, errStatus, err := saveCompanyDomain.Handle(uint(1), "Company test name", &description, &website, &linkedin, &glassdoor, &instagram)
		assert.Nil(t, createdCompany)
		assert.Equal(t, errStatus, 400)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockCompanyRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
	})

	t.Run("Should updated a company when pass an valid ID", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		mockCompanyRepository.On(
			"CreateOrUpdate",
			uint(1),
			"Company test name",
			&description,
			&website,
			&linkedin,
			&glassdoor,
			&instagram,
		).Return(1, nil)
		saveCompanyDomain := company_domain.SaveCompany{CompanyRepository: mockCompanyRepository}
		createdCompany, errStatus, err := saveCompanyDomain.Handle(1, "Company test name", &description, &website, &linkedin, &glassdoor, &instagram)
		assert.Equal(t, createdCompany, expectedCompany)
		assert.Equal(t, errStatus, 0)
		assert.Nil(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 2)
	})

	t.Run("Should return company created when ID passed is equal zero", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		mockCompanyRepository.On(
			"CreateOrUpdate",
			uint(0),
			"Company test name",
			&description,
			&website,
			&linkedin,
			&glassdoor,
			&instagram,
		).Return(1, nil)
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		saveCompanyDomain := company_domain.SaveCompany{CompanyRepository: mockCompanyRepository}
		createdCompany, errStatus, err := saveCompanyDomain.Handle(0, "Company test name", &description, &website, &linkedin, &glassdoor, &instagram)
		assert.Equal(t, createdCompany, expectedCompany)
		assert.Equal(t, errStatus, 0)
		assert.Nil(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockCompanyRepository.AssertCalled(t, "GetById", uint(1))
	})
}
