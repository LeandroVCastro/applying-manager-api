package company_repository_unit_test

import (
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/company"
	"github.com/stretchr/testify/mock"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) GetById(id uint) *company_repository.SelectNoRelations {
	args := m.Called(id)
	return args.Get(0).(*company_repository.SelectNoRelations)
}

func (m *MockCompanyRepository) CreateOrUpdate(id uint, name string, description, website, linkedin, glassdoor, instagram *string) (uint, error) {
	args := m.Called(id, name, description, website, linkedin, glassdoor, instagram)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockCompanyRepository) ListAll() (listedCompanies []*company_repository.SelectNoRelations, err error) {
	args := m.Called()
	return args.Get(0).([]*company_repository.SelectNoRelations), args.Error(1)
}

func (m *MockCompanyRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
