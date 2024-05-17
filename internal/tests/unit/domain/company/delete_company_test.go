package company_domain

import (
	"errors"
	"testing"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/company"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCompanyDomain(t *testing.T) {
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

	t.Run("Should return error when an ID is not provided", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		deleteCompanyDomain := company_domain.DeleteCompany{CompanyRepository: mockCompanyRepository}
		errStatus, err := deleteCompanyDomain.Handle(0)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 404 when company not found", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		var expectedNilCompany *entity.Company
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedNilCompany)
		deleteCompanyDomain := company_domain.DeleteCompany{CompanyRepository: mockCompanyRepository}
		errStatus, err := deleteCompanyDomain.Handle(uint(1))
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return error 500 when something went to Delete repository method", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		mockCompanyRepository.On("Delete", uint(1)).Return(errors.New("error"))
		deleteCompanyDomain := company_domain.DeleteCompany{CompanyRepository: mockCompanyRepository}
		errStatus, err := deleteCompanyDomain.Handle(uint(1))
		assert.Equal(t, 500, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockCompanyRepository.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("Should return error nil when company is deleted successfully", func(t *testing.T) {
		mockCompanyRepository := new(company_repository.MockCompanyRepository)
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		mockCompanyRepository.On("Delete", uint(1)).Return(nil)
		deleteCompanyDomain := company_domain.DeleteCompany{CompanyRepository: mockCompanyRepository}
		errStatus, err := deleteCompanyDomain.Handle(uint(1))
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockCompanyRepository.AssertNumberOfCalls(t, "Delete", 1)
	})
}
