package platform_repository

import (
	"errors"
	"fmt"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type PlatformRepositoryInterface interface {
	GetById(id uint) *entity.Platform
	CreateOrUpdate(id uint, name string, website *string) (savedId uint, err error)
	ListAll() (platforms []*entity.Platform, err error)
}

type PlatformRepository struct {
	connection *gorm.DB
}

func (repository PlatformRepository) ListAll() (listedPlatforms []*entity.Platform, err error) {
	result := repository.connection.Order("id ASC").Find(&listedPlatforms)
	if result.Error != nil {
		err = errors.New(result.Error.Error())
		return
	}
	return
}

func (repository PlatformRepository) GetById(id uint) *entity.Platform {
	var platform = entity.Platform{}
	result := repository.connection.First(&platform, id)
	if result.Error != nil {
		return nil
	}
	return &platform
}

func (repository PlatformRepository) CreateOrUpdate(
	id uint,
	name string,
	website *string,
) (savedId uint, err error) {
	platformParams := entity.Platform{Name: name}
	if *website != "" {
		platformParams.Website = website
	}
	var result *gorm.DB
	if id != 0 {
		platformParams.ID = id
		result = repository.connection.Updates(&platformParams)
	} else {
		result = repository.connection.Create(&platformParams)
	}
	if result.Error != nil {
		fmt.Println("Error on createOrUpdate PlatformRepository: " + result.Error.Error())
		err = result.Error
		return
	}
	savedId = platformParams.ID
	return
}

func PlatformRepositoryFactory() PlatformRepository {
	repository := PlatformRepository{
		connection: configs.Connection,
	}
	return repository
}
