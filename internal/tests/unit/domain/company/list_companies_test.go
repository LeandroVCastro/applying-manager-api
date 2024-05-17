package company_domain_unit_test

import (
	"errors"
	"testing"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	company_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/company"
	"github.com/stretchr/testify/assert"
)

func TestListCompaniesDomain(t *testing.T) {
	t.Run("Should return exactly what repository returns", func(t *testing.T) {
		mockCompanyRepository := new(company_repository_unit_test.MockCompanyRepository)
		expectedCompanies := []*entity.Company{}
		expectedCompanies = append(expectedCompanies, &entity.Company{
			ID:   1,
			Name: "Company test name",
		}, &entity.Company{
			ID:   2,
			Name: "Another company test name",
		})
		mockCompanyRepository.On("ListAll").Return(expectedCompanies, nil)
		listCompanyDomain := company_domain.ListCompanies{CompanyRepository: mockCompanyRepository}
		listedCompanies, errStatus, err := listCompanyDomain.Handle()
		assert.Equal(t, expectedCompanies, listedCompanies)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		mockCompanyRepository := new(company_repository_unit_test.MockCompanyRepository)
		expectedCompanies := []*entity.Company{}
		mockCompanyRepository.On("ListAll").Return(expectedCompanies, errors.New("Error to select companies"))
		listCompanyDomain := company_domain.ListCompanies{CompanyRepository: mockCompanyRepository}
		listedCompanies, errStatus, err := listCompanyDomain.Handle()
		assert.Equal(t, expectedCompanies, listedCompanies)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})
}
