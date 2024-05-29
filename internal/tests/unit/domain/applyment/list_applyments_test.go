package applyment_domain_unit_test

import (
	"errors"
	"testing"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/applyment"
	"github.com/stretchr/testify/assert"
)

func TestListCompaniesDomain(t *testing.T) {
	t.Run("Should return exactly what repository returns", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		expectedApplyments := []*entity.Applyment{}
		expectedApplyments = append(expectedApplyments, &entity.Applyment{
			ID:    1,
			Title: "Applyment test name",
		}, &entity.Applyment{
			ID:    2,
			Title: "Another applyment test name",
		})
		mockApplymentRepository.On("ListAll").Return(expectedApplyments, nil)
		listApplymentsDomain := applyment_domain.ListApplyments{ApplymentRepository: mockApplymentRepository}
		listedApplyments, errStatus, err := listApplymentsDomain.Handle()
		assert.Equal(t, expectedApplyments, listedApplyments)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		expectedApplyments := []*entity.Applyment{}
		mockApplymentRepository.On("ListAll").Return(expectedApplyments, errors.New("Error to select companies"))
		listApplymentsDomain := applyment_domain.ListApplyments{ApplymentRepository: mockApplymentRepository}
		listedApplyments, errStatus, err := listApplymentsDomain.Handle()
		assert.Equal(t, expectedApplyments, listedApplyments)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})
}
