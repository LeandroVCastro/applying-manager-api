package applyment_domain_unit_test

import (
	"testing"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/applyment"

	"github.com/stretchr/testify/assert"
)

func TestGetPlatformDomain(t *testing.T) {
	// var website string = "testewebsite"
	var expectedApplyment = &entity.Applyment{
		ID:    1,
		Title: "Test applyment",
		// Name:    "Platform test name",
		// Website: &website,
	}

	t.Run("Should return error 400 when ID is not provided", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		getApplymentDomain := applyment_domain.GetApplyment{ApplymentRepository: mockApplymentRepository}
		applyment, errStatus, err := getApplymentDomain.Handle(0)
		assert.Nil(t, applyment)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "should be provided an ID", err.Error())
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 404 when applyment was not found", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		var foundApplyment *entity.Applyment
		mockApplymentRepository.On("GetById", uint(1)).Return(foundApplyment)
		getApplymentDomain := applyment_domain.GetApplyment{ApplymentRepository: mockApplymentRepository}
		applyment, errStatus, err := getApplymentDomain.Handle(uint(1))
		assert.Nil(t, applyment)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "applyment not found", err.Error())
		mockApplymentRepository.AssertCalled(t, "GetById", uint(1))
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return exactly what repository returns", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		getApplymentDomain := applyment_domain.GetApplyment{ApplymentRepository: mockApplymentRepository}
		applyment, errStatus, err := getApplymentDomain.Handle(uint(1))
		assert.Equal(t, expectedApplyment, applyment)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertCalled(t, "GetById", uint(1))
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
	})
}
