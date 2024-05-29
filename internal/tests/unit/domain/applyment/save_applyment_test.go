package applyment_domain_unit_test

import (
	"errors"
	"testing"
	"time"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/applyment"
	company_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/company"
	platform_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/platform"

	"github.com/stretchr/testify/assert"
)

func getMocks() (
	*applyment_repository_unit_test.MockApplymentRepository,
	*company_repository_unit_test.MockCompanyRepository,
	*platform_repository_unit_test.MockPlatformRepository,
) {
	mockApplymentRepository := new(applyment_repository_unit_test.MockApplymentRepository)
	mockCompanyRepository := new(company_repository_unit_test.MockCompanyRepository)
	mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
	return mockApplymentRepository, mockCompanyRepository, mockPlatformRepository
}

func getDomain(
	mock_applyment_repository *applyment_repository_unit_test.MockApplymentRepository,
	mock_company_repository *company_repository_unit_test.MockCompanyRepository,
	mock_platform_repository *platform_repository_unit_test.MockPlatformRepository,
) applyment_domain.SaveApplyment {
	saveApplymentDomain := applyment_domain.SaveApplyment{
		ApplymentRepository: mock_applyment_repository,
		CompanyRepository:   mock_company_repository,
		PlatformRepository:  mock_platform_repository,
	}
	return saveApplymentDomain
}

func TestSaveApplymentDomain(t *testing.T) {
	var description string = "teste"
	var link string = "testewebsite"

	var applied_at time.Time
	var expectedApplyment = &entity.Applyment{
		ID:    1,
		Title: "Applyment test name",
	}
	t.Run("Should return error 404 when platform is not found", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		var expectedPlatform *entity.Platform
		var company_id uint
		var platform_id uint = 1
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			1,
			"name teste",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "platform not found", err.Error())
	})

	t.Run("Should return error 404 when company is not found", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		var expectedCompany *entity.Company
		var company_id uint = 1
		var platform_id uint
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			1,
			"name teste",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "company not found", err.Error())
	})

	t.Run("Should return error 404 when an ID is provided and applyment is not found", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		var expectedApplyment *entity.Applyment
		var company_id uint
		var platform_id uint
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(1, "name teste", &description, &link, &company_id, &platform_id, &applied_at)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return error 400 when update fails", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		var company_id uint
		var platform_id uint
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
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			uint(1),
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
	})

	t.Run("Should updated an applyment when pass an valid ID", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		mockApplymentRepository.On("GetById", uint(1)).Return(expectedApplyment)
		var company_id uint
		var platform_id uint
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
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			1,
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Equal(t, expectedApplyment, createdApplyment)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 2)
	})

	t.Run("Should return applyment created when ID passed is equal zero", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		var company_id uint
		var platform_id uint
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
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			0,
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Equal(t, expectedApplyment, createdApplyment)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertCalled(t, "GetById", uint(1))
	})

	t.Run("Should return applyment created when ID passed is equal zero and get company too", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		var company_id uint = 1
		var platform_id uint = 1
		var expectedCompany = &entity.Company{
			ID:   1,
			Name: "Company test",
		}
		var expectedPlatform = &entity.Platform{
			ID:   1,
			Name: "Platform test",
		}
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
		mockCompanyRepository.On("GetById", uint(1)).Return(expectedCompany)
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			0,
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Equal(t, expectedApplyment, createdApplyment)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockApplymentRepository.AssertCalled(t, "GetById", uint(1))
	})

	t.Run("Should return error 400 when Create fails", func(t *testing.T) {
		mockApplymentRepository, mockCompanyRepository, mockPlatformRepository := getMocks()
		var company_id uint
		var platform_id uint
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
		domain := getDomain(mockApplymentRepository, mockCompanyRepository, mockPlatformRepository)
		createdApplyment, errStatus, err := domain.Handle(
			0,
			"Applyment test name",
			&description,
			&link,
			&company_id,
			&platform_id,
			&applied_at,
		)
		assert.Nil(t, createdApplyment)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockApplymentRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockApplymentRepository.AssertNumberOfCalls(t, "GetById", 0)
	})
}
