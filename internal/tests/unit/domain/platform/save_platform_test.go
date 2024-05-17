package platform_domain_unit_test

import (
	"errors"
	"testing"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	platform_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/platform"

	"github.com/stretchr/testify/assert"
)

func TestSaveCompanyDomain(t *testing.T) {
	var website string = "testewebsite"
	var expectedPlatform = &entity.Platform{
		ID:      1,
		Name:    "Platform test name",
		Website: &website,
	}
	t.Run("Should return error 404 when an ID is provided and platform is not found", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		var foundPlatform *entity.Platform
		mockPlatformRepository.On("GetById", uint(1)).Return(foundPlatform)
		savePlatformDomain := platform_domain.SavePlatform{PlatformRepository: mockPlatformRepository}
		createdCompany, errStatus, err := savePlatformDomain.Handle(1, "name teste", nil)
		assert.Nil(t, createdCompany)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
	})

	t.Run("Should create a new platform when an ID is not provided", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockPlatformRepository.On("CreateOrUpdate", uint(0), "Platform test name", &website).Return(1, nil)
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		savePlatform := platform_domain.SavePlatform{PlatformRepository: mockPlatformRepository}
		createdPlatform, errStatus, err := savePlatform.Handle(0, "Platform test name", &website)
		assert.Equal(t, expectedPlatform, createdPlatform)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockPlatformRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
	})

	t.Run("Should return error 400 when Create fails", func(t *testing.T) {
		mockCompanyRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockCompanyRepository.On(
			"CreateOrUpdate",
			uint(0),
			"Platform test name",
			&website,
		).Return(0, errors.New("error on create"))
		savePlatformDomain := platform_domain.SavePlatform{PlatformRepository: mockCompanyRepository}
		createdPlatform, errStatus, err := savePlatformDomain.Handle(0, "Platform test name", &website)
		assert.Nil(t, createdPlatform)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockCompanyRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockCompanyRepository.AssertNumberOfCalls(t, "GetById", 0)
	})
}
