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
