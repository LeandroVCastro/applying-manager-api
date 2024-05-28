package applyment_domain_unit_test

import (
	"errors"
	"testing"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/applyment"
	"github.com/stretchr/testify/assert"
)

func TestDeleteApplymentDomain(t *testing.T) {
	var expectedApplyment = &entity.Applyment{
		ID:    1,
		Title: "Applyment test name",
	}

	t.Run("Should return error when an ID is not provided", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		deleteApplymentDomain := applyment_domain.DeleteApplyment{ApplymentRepository: mockApplymentRepository}
		errStatus, err := deleteApplymentDomain.Handle(0)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 404 when applyment not found", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		var expectedNilCompany *entity.Applyment
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedNilCompany)
		deleteApplymentDomain := applyment_domain.DeleteApplyment{ApplymentRepository: mockApplymentRepository}
		errStatus, err := deleteApplymentDomain.Handle(uint(1))
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return error 500 when something went wrong to Delete repository method", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		mockApplymentRepository.On("Delete", uint(1)).Return(errors.New("error"))
		deleteApplymentDomain := applyment_domain.DeleteApplyment{ApplymentRepository: mockApplymentRepository}
		errStatus, err := deleteApplymentDomain.Handle(uint(1))
		assert.Equal(t, 500, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("Should return error nil when company is deleted successfully", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		mockApplymentRepository.On("Delete", uint(1)).Return(nil)
		deleteCompanyDomain := applyment_domain.DeleteApplyment{ApplymentRepository: mockApplymentRepository}
		errStatus, err := deleteCompanyDomain.Handle(uint(1))
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "Delete", 1)
	})
}
