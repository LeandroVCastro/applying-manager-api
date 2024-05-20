package platform_domain_unit_test

import (
	"errors"
	"testing"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	platform_repository_unit_test "github.com/LeandroVCastro/applying-manager-api/internal/tests/unit/repository/platform"
	"github.com/stretchr/testify/assert"
)

func TestListPlatformsDomain(t *testing.T) {
	t.Run("Should return exactly what repository returns", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		expectedPlatforms := []*entity.Platform{}
		expectedPlatforms = append(expectedPlatforms, &entity.Platform{
			ID:   1,
			Name: "Platform test name",
		}, &entity.Platform{
			ID:   2,
			Name: "Another platform test name",
		})
		mockPlatformRepository.On("ListAll").Return(expectedPlatforms, nil)
		listPlatformDomain := platform_domain.ListPlatforms{PlatformRepository: mockPlatformRepository}
		listedCompanies, errStatus, err := listPlatformDomain.Handle()
		assert.Equal(t, expectedPlatforms, listedCompanies)
		assert.Equal(t, 0, errStatus)
		assert.Nil(t, err)
		mockPlatformRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		mockPlatformRepository := new(platform_repository_unit_test.MockPlatformRepository)
		expectedPlatforms := []*entity.Platform{}
		mockPlatformRepository.On("ListAll").Return(expectedPlatforms, errors.New("Error to select platforms"))
		listPlatformDomain := platform_domain.ListPlatforms{PlatformRepository: mockPlatformRepository}
		listedPlatforms, errStatus, err := listPlatformDomain.Handle()
		assert.Equal(t, expectedPlatforms, listedPlatforms)
		assert.Equal(t, 400, errStatus)
		assert.Error(t, err)
		mockPlatformRepository.AssertNumberOfCalls(t, "ListAll", 1)
	})
}
