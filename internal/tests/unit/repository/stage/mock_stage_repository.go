package stage_repository_unit_test

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockStageRepository struct {
	mock.Mock
}

func (m *MockStageRepository) ListAll() (listedStages []*entity.Stage, err error) {
	args := m.Called()
	return args.Get(0).([]*entity.Stage), args.Error(1)
}
