package platform_repository_unit_test

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockPlatformRepository struct {
	mock.Mock
}

func (m *MockPlatformRepository) GetById(id uint) *entity.Platform {
	args := m.Called(id)
	return args.Get(0).(*entity.Platform)
}

func (m *MockPlatformRepository) CreateOrUpdate(id uint, name string, website *string) (savedId uint, err error) {
	args := m.Called(id, name, website)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockPlatformRepository) ListAll() (listedPlatforms []*entity.Platform, err error) {
	args := m.Called()
	return args.Get(0).([]*entity.Platform), args.Error((1))
	// return uint(args.Int(0)), args.Error(1)
}
