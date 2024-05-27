package platform_domain_unit_test

import (
	"errors"
	"testing"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
	platform_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/platform"

	"github.com/stretchr/testify/assert"
)

func TestSavePlatformDomain(t *testing.T) {
	var website string = "testewebsite"
	var expectedPlatform = &platform_repository.SelectNoRelations{
		ID:      1,
		Name:    "Platform test name",
		Website: &website,
	}
	t.Run("Should return error 404 when an ID is provided and platform is not found", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		var foundPlatform *platform_repository.SelectNoRelations
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
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockPlatformRepository.On(
			"CreateOrUpdate",
			uint(0),
			"Platform test name",
			&website,
		).Return(0, errors.New("error on create"))
		savePlatformDomain := platform_domain.SavePlatform{PlatformRepository: mockPlatformRepository}
		createdPlatform, errStatus, err := savePlatformDomain.Handle(0, "Platform test name", &website)
		assert.Nil(t, createdPlatform)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockPlatformRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 400 when Update fails", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		mockPlatformRepository.On(
			"CreateOrUpdate",
			uint(1),
			"Platform test name updated",
			&website,
		).Return(1, errors.New("error on update"))
		savePlatformDomain := platform_domain.SavePlatform{PlatformRepository: mockPlatformRepository}
		savedPlatform, errStatus, err := savePlatformDomain.Handle(uint(1), "Platform test name updated", &website)
		assert.Nil(t, savedPlatform)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "error on update platform", err.Error())
		mockPlatformRepository.AssertNumberOfCalls(t, "CreateOrUpdate", 1)
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
	})
}
