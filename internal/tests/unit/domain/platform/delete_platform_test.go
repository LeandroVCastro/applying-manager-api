package platform_domain_unit_test

import (
	"errors"
	"testing"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	platform_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/platform"

	"github.com/stretchr/testify/assert"
)

func TestDeletePlatformDomain(t *testing.T) {
	var website string = "testewebsite"
	var expectedPlatform = &entity.Platform{
		ID:      1,
		Name:    "Platform test name",
		Website: &website,
	}

	t.Run("Should return error 400 when ID is not provided", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		deletePlatformDomain := platform_domain.DeletePlatform{PlatformRepository: mockPlatformRepository}
		errStatus, err := deletePlatformDomain.Handle(0)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "an ID should be provided", err.Error())
		mockPlatformRepository.AssertNumberOfCalls(t, "Delete", 0)
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 0)
	})

	t.Run("Should return error 404 when platform was not found", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		var foundPlatform *entity.Platform
		mockPlatformRepository.On("GetById", uint(1)).Return(foundPlatform)
		deletePlatformDomain := platform_domain.DeletePlatform{PlatformRepository: mockPlatformRepository}
		errStatus, err := deletePlatformDomain.Handle(uint(1))
		assert.Equal(t, 404, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "platform not found", err.Error())
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockPlatformRepository.AssertNumberOfCalls(t, "Delete", 0)
	})

	t.Run("Should return error 500 when repository Delete fails", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		mockPlatformRepository.On("Delete", uint(1)).Return(errors.New("server error"))
		deletePlatformDomain := platform_domain.DeletePlatform{PlatformRepository: mockPlatformRepository}
		errStatus, err := deletePlatformDomain.Handle(uint(1))
		assert.Equal(t, 500, errStatus)
		assert.Error(t, err)
		assert.Equal(t, "error on delete statement", err.Error())
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockPlatformRepository.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("Should execute Delete and not return error when all is okay", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		mockPlatformRepository.On("GetById", uint(1)).Return(expectedPlatform)
		mockPlatformRepository.On("Delete", uint(1)).Return(nil)
		deletePlatformDomain := platform_domain.DeletePlatform{PlatformRepository: mockPlatformRepository}
		errStatus, err := deletePlatformDomain.Handle(uint(1))
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockPlatformRepository.AssertCalled(t, "GetById", uint(1))
		mockPlatformRepository.AssertNumberOfCalls(t, "GetById", 1)
		mockPlatformRepository.AssertNumberOfCalls(t, "Delete", 1)
	})
}
