package platform_domain

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
)

type ListPlatforms struct {
	PlatformRepository platform_repository.PlatformRepositoryInterface
}

func (p ListPlatforms) Handle() (platforms []*entity.Platform, errStatus int, err error) {
	platforms, err = p.PlatformRepository.ListAll()
	if err != nil {
		errStatus = 400
	}
	return
}

func ListPlatformsFactory() ListPlatforms {
	return ListPlatforms{
		PlatformRepository: platform_repository.PlatformRepositoryFactory(),
	}
}
