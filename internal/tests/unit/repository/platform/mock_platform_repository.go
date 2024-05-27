package platform_repository_unit_test

import (
	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
	"github.com/stretchr/testify/mock"
)

type MockPlatformRepository struct {
	mock.Mock
}

func (m *MockPlatformRepository) GetById(id uint) *platform_repository.SelectNoRelations {
	args := m.Called(id)
	return args.Get(0).(*platform_repository.SelectNoRelations)
}

func (m *MockPlatformRepository) CreateOrUpdate(id uint, name string, website *string) (savedId uint, err error) {
	args := m.Called(id, name, website)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockPlatformRepository) ListAll() (listedPlatforms []*platform_repository.SelectNoRelations, err error) {
	args := m.Called()
	return args.Get(0).([]*platform_repository.SelectNoRelations), args.Error(1)
}

func (m *MockPlatformRepository) Delete(id uint) (err error) {
	args := m.Called(id)
	return args.Error(0)
}
