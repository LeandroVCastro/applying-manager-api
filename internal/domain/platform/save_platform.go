package platform_domain

import (
	"errors"

	"strconv"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
)

type SavePlatform struct {
	PlatformRepository platform_repository.PlatformRepositoryInterface
}

func (c SavePlatform) Handle(
	id uint,
	name string,
	website *string,
) (savedPlatform *entity.Platform, errStatus int, err error) {
	if id != 0 {
		if platform := c.PlatformRepository.GetById(id); platform == nil {
			err = errors.New("platform not found by ID '" + strconv.FormatUint(uint64(id), 10) + "'")
			errStatus = 404
			return
		}
		// savedId, updatedErr := c.CompanyRepository.CreateOrUpdate(id, name, description, website, linkedin, glassdoor, instagram)
		// if updatedErr != nil {
		// 	err = errors.New("error on update company")
		// 	errStatus = 400
		// 	return
		// }
		// savedCompany = c.CompanyRepository.GetById(savedId)
		// return
	}
	savedId, saveErr := c.PlatformRepository.CreateOrUpdate(0, name, website)
	if saveErr != nil {
		err = errors.New("error on create platform")
		errStatus = 400
		return
	}
	savedPlatform = c.PlatformRepository.GetById(savedId)
	return
}

func SavePlatformFactory() SavePlatform {
	return SavePlatform{
		PlatformRepository: platform_repository.PlatformRepositoryFactory(),
	}
}
