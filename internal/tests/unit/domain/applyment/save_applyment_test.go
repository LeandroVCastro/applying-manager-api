package applyment_domain_unit_test

import (
	"errors"
	"testing"
	"time"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/applyment"

	"github.com/stretchr/testify/assert"
)

func TestSaveApplymentDomain(t *testing.T) {
	var description string = "teste"
	var link string = "testewebsite"
	var company_id uint = 1
	var platform_id uint = 1
	var applied_at time.Time
	var expectedApplyment = &entity.Applyment{
		ID:    1,
		Title: "Applyment test name",
	}
	t.Run("Should return error 404 when an ID is provided and applyment is not found", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		var expectedCompany *entity.Applyment
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedCompany)
		saveApplymentDomain := applyment_domain.SaveApplyment{ApplymentRepository: mockApplymentRepository}
		createdApplyment, errStatus, err := saveApplymentDomain.Handle(1, "name teste", &description, &link, &company_id, &platform_id, &applied_at)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return error 400 when update fails", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		mockApplymentRepository.On(
			"CreateOrUpdate",
			uint(1),
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		).Return(0, errors.New("error on update"))
		saveApplymentDomain := applyment_domain.SaveApplyment{ApplymentRepository: mockApplymentRepository}
		createdApplyment, errStatus, err := saveApplymentDomain.Handle(uint(1), "Applyment test name", &description, &link, &company_id, &platform_id, &applied_at)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
	})

	t.Run("Should updated an applyment when pass an valid ID", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		mockApplymentRepository.On(
			"CreateOrUpdate",
			uint(1),
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		).Return(1, nil)
		saveApplymentDomain := applyment_domain.SaveApplyment{ApplymentRepository: mockApplymentRepository}
		createdApplyment, errStatus, err := saveApplymentDomain.Handle(1, "Applyment test name", &description, &link, &company_id, &platform_id, &applied_at)
		assert.Equal(t, expectedApplyment, createdApplyment)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 2)
	})

	t.Run("Should return applyment created when ID passed is equal zero", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On(
			"CreateOrUpdate",
			uint(0),
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		).Return(1, nil)
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		saveApplymentDomain := applyment_domain.SaveApplyment{ApplymentRepository: mockApplymentRepository}
		createdApplyment, errStatus, err := saveApplymentDomain.Handle(0, "Applyment test name", &description, &link, &company_id, &platform_id, &applied_at)
		assert.Equal(t, expectedApplyment, createdApplyment)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertCalled(t, "GetById", uint(1))
	})

	t.Run("Should return error 400 when Create fails", func(t *testing.T) {
		mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
		mockApplymentRepository.On(
			"CreateOrUpdate",
			uint(0),
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		).Return(0, errors.New("error on create"))
		saveApplymentDomain := applyment_domain.SaveApplyment{ApplymentRepository: mockApplymentRepository}
		createdApplyment, errStatus, err := saveApplymentDomain.Handle(0, "Applyment test name", &description, &link, &company_id, &platform_id, &applied_at)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 0)
	})
}
