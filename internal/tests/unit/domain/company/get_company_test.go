package company_domain_unit_test

import (
	"testing"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/company"
	company_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/company"
	"github.com/stretchr/testify/assert"
)

func TestGetCompanyDomain(t *testing.T) {
	var description string = "teste"
	var website string = "testewebsite"
	var linkedin string = "testelinkedin"
	var glassdoor string = "testeglassdoor"
	var instagram string = "testeinstagram"
	var expectedCompany = &company_repository.SelectNoRelations{
		ID:          1,
		Name:        "Company test name",
		Description: &description,
		Website:     &website,
		Linkedin:    &linkedin,
		Glassdoor:   &glassdoor,
		Instagram:   &instagram,
	}

	t.Run("Should return error when an ID is not provided", func(t *testing.T) {
		mockCompanyRepository := new(company_repository_unit_test.MockCompanyRepository)
		listCompanyDomain := company_domain.GetCompany{CompanyRepository: mockCompanyRepository}
		_, errStatus, err := listCompanyDomain.Handle(0)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 404 when company not found", func(t *testing.T) {
		mockCompanyRepository := new(company_repository_unit_test.MockCompanyRepository)
		var expectedCompany *company_repository.SelectNoRelations
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		listCompanyDomain := company_domain.GetCompany{CompanyRepository: mockCompanyRepository}
		_, errStatus, err := listCompanyDomain.Handle(uint(1))
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return the company returned by repository", func(t *testing.T) {
		mockCompanyRepository := new(company_repository_unit_test.MockCompanyRepository)
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		listCompanyDomain := company_domain.GetCompany{CompanyRepository: mockCompanyRepository}
		company, errStatus, err := listCompanyDomain.Handle(uint(1))
		assert.Equal(t, expectedCompany, company)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 1)
	})
}
