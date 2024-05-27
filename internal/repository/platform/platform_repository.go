package platform_repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type PlatformRepositoryInterface interface {
	GetById(id uint) *SelectNoRelations
	CreateOrUpdate(id uint, name string, website *string) (savedId uint, err error)
	ListAll() (platforms []*SelectNoRelations, err error)
	Delete(id uint) error
}

type PlatformRepository struct {
	connection *gorm.DB
}

type SelectNoRelations struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Website   *string        `json:"website"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (repository PlatformRepository) ListAll() (listedPlatforms []*SelectNoRelations, err error) {
	// result := repository.connection.Order("id ASC").Find(&listedPlatforms)
	result := repository.connection.Table("platforms").Order("id ASC").Find(&listedPlatforms)
	if result.Error != nil {
		err = errors.New(result.Error.Error())
		return
	}
	return
}

func (repository PlatformRepository) GetById(id uint) *SelectNoRelations {
	var platform = SelectNoRelations{}
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

func (repository PlatformRepository) Delete(id uint) error {
	var platform = entity.Platform{}
	result := repository.connection.Where("ID = ?", id).Delete(&platform)
	if result.Error != nil {
		err := errors.New(result.Error.Error())
		return err
	}
	return nil
}
