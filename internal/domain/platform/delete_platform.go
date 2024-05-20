package platform_domain

import (
	"errors"

	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
)

type DeletePlatform struct {
	PlatformRepository platform_repository.PlatformRepositoryInterface
}

func (p DeletePlatform) Handle(id uint) (errStatus int, err error) {
	if id == 0 {
		err = errors.New("an ID should be provided")
		errStatus = 400
		return
	}
	platform := p.PlatformRepository.GetById(id)
	if platform == nil {
		err = errors.New("platform not found")
		errStatus = 404
		return
	}
	errDelete := p.PlatformRepository.Delete(id)
	if errDelete != nil {
		err = errors.New("error on delete statement")
		errStatus = 500
		return
	}
	return
}

func DeletePlatformFactory() DeletePlatform {
	return DeletePlatform{
		PlatformRepository: platform_repository.PlatformRepositoryFactory(),
	}
}
