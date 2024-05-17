package company_repository

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) GetById(id uint) *entity.Company {
	args := m.Called(id)
	return args.Get(0).(*entity.Company)
}

func (m *MockCompanyRepository) CreateOrUpdate(id uint, name string, description, website, linkedin, glassdoor, instagram *string) (uint, error) {
	args := m.Called(id, name, description, website, linkedin, glassdoor, instagram)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockCompanyRepository) ListAll() (listedCompanies []*entity.Company, err error) {
	args := m.Called()
	return args.Get(0).([]*entity.Company), args.Error(1)
}

func (m *MockCompanyRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
