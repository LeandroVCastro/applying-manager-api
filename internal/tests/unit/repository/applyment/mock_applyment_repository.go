package applyment_repository_unit_test

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockApplymentRepository struct {
	mock.Mock
}

func (m *MockApplymentRepository) GetById(id uint) *entity.Applyment {
	args := m.Called(id)
	return args.Get(0).(*entity.Applyment)
}

func (m *MockApplymentRepository) ListAll() (applyments []*entity.Applyment, err error) {
	args := m.Called()
	return args.Get(0).([]*entity.Applyment), args.Error(1)
}

func (m *MockApplymentRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
