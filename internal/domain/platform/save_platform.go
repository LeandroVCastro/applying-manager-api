package platform_domain

import (
	"errors"

	"strconv"

	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
)

type SavePlatform struct {
	PlatformRepository platform_repository.PlatformRepositoryInterface
}

func (p SavePlatform) Handle(
	id uint,
	name string,
	website *string,
) (savedPlatform *platform_repository.SelectNoRelations, errStatus int, err error) {
	if id != 0 {
		if platform := p.PlatformRepository.GetById(id); platform == nil {
			err = errors.New("platform not found by ID '" + strconv.FormatUint(uint64(id), 10) + "'")
			errStatus = 404
			return
		}
		savedId, updatedErr := p.PlatformRepository.CreateOrUpdate(id, name, website)
		if updatedErr != nil {
			err = errors.New("error on update platform")
			errStatus = 400
			return
		}
		savedPlatform = p.PlatformRepository.GetById(savedId)
		return
	}
	savedId, saveErr := p.PlatformRepository.CreateOrUpdate(0, name, website)
	if saveErr != nil {
		err = errors.New("error on create platform")
		errStatus = 400
		return
	}
	savedPlatform = p.PlatformRepository.GetById(savedId)
	return
}

func SavePlatformFactory() SavePlatform {
	return SavePlatform{
		PlatformRepository: platform_repository.PlatformRepositoryFactory(),
	}
}
