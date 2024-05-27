package platform_domain

import (
	"errors"

	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
)

type GetPlatform struct {
	PlatformRepository platform_repository.PlatformRepositoryInterface
}

func (p GetPlatform) Handle(id uint) (platform *platform_repository.SelectNoRelations, errStatus int, err error) {
	if id == 0 {
		err = errors.New("should be provided an ID")
		errStatus = 400
		return
	}
	platform = p.PlatformRepository.GetById(id)
	if platform == nil {
		err = errors.New("platform not found")
		errStatus = 404
		return
	}
	return

}

func GetPlatformFactory() GetPlatform {
	return GetPlatform{
		PlatformRepository: platform_repository.PlatformRepositoryFactory(),
	}
}
