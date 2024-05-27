package platform_domain_unit_test

import (
	"testing"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
	platform_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/platform"

	"github.com/stretchr/testify/assert"
)

func TestGetPlatformDomain(t *testing.T) {
	var website string = "testewebsite"
	var expectedPlatform = &platform_repository.SelectNoRelations{
		ID:      1,
		Name:    "Platform test name",
		Website: &website,
	}

	t.Run("Should return error 400 when ID is not provided", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		getPlatformDomain := platform_domain.GetPlatform{PlatformRepository: mockPlatformRepository}
		platform, errStatus, err := getPlatformDomain.Handle(0)
		assert.Nil(t, platform)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "should be provided an ID", err.Error())
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 404 when platform was not found", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		var foundPlatform *platform_repository.SelectNoRelations
		mockPlatformRepository.On("GetById", uint(1)).Return(foundPlatform)
		getPlatformDomain := platform_domain.GetPlatform{PlatformRepository: mockPlatformRepository}
		platform, errStatus, err := getPlatformDomain.Handle(uint(1))
		assert.Nil(t, platform)
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "platform not found", err.Error())
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
	})

	t.Run("Should return exactly what repository returns", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		getPlatformDomain := platform_domain.GetPlatform{PlatformRepository: mockPlatformRepository}
		platform, errStatus, err := getPlatformDomain.Handle(uint(1))
		assert.Equal(t, expectedPlatform, platform)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
	})
}
