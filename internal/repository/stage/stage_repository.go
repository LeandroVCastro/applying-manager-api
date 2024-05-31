package stage_repository

import (
	"errors"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type StageRepositoryInterface interface {
	ListAll() (applyments []*entity.Stage, err error)
}

type StageRepository struct {
	connection *gorm.DB
}

func (repository StageRepository) ListAll() (stages []*entity.Stage, err error) {
	result := repository.connection.Select([]string{"id", "title", "description", "created_at", "updated_at"}).Order("id ASC").Find(&stages)
	if result.Error != nil {
		err = errors.New(result.Error.Error())
		return
	}
	return
}

func StageRepositoryFactory() StageRepository {
	return StageRepository{
		connection: configs.Connection,
	}
}
