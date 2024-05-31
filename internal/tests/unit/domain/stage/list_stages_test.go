package stage_domain_unit_test

import (
	"errors"
	"testing"

	stage_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/stage"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	stage_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/stage"
	"github.com/stretchr/testify/assert"
)

func getMocks() *stage_repository_unit_test.MockStageRepository {
	mockStageRepository := new(stage_repository_unit_test.MockStageRepository)
	return mockStageRepository
}

func getDomain(
	mockStageRepository *stage_repository_unit_test.MockStageRepository,
) stage_domain.ListStages {
	listStagesDomain := stage_domain.ListStages{StageRepository: mockStageRepository}
	return listStagesDomain
}

func TestListStagesDomain(t *testing.T) {
	t.Run("Should return exactly what repository returns", func(t *testing.T) {
		mockStageRepository := getMocks()
		expectedStages := []*entity.Stage{}
		expectedStages = append(expectedStages, &entity.Stage{
			ID:          1,
			Title:       "Stage test name",
			Description: "Stage description test",
		}, &entity.Stage{
			ID:          2,
			Title:       "Stage test name 2",
			Description: "Stage description test 2",
		})
		mockStageRepository.On("ListAll").Return(expectedStages, nil)
		domain := getDomain(mockStageRepository)
		listedStages, errStatus, err := domain.Handle()
		assert.Equal(t, expectedStages, listedStages)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockStageRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		mockStageRepository := getMocks()
		expectedStages := []*entity.Stage{}
		mockStageRepository.On("ListAll").Return(expectedStages, errors.New("Error to select stages"))
		domain := getDomain(mockStageRepository)
		listedStages, errStatus, err := domain.Handle()
		assert.Equal(t, expectedStages, listedStages)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockStageRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})
}
