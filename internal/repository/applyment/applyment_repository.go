package applyment_repository

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type ApplymentRepositoryInterface interface {
	GetById(id uint) *entity.Applyment
}

type ApplymentRepository struct {
	connection *gorm.DB
}

func (repository ApplymentRepository) GetById(id uint) *entity.Applyment {
	var applyment = entity.Applyment{}
	result := repository.connection.First(&applyment, id)
	if result.Error != nil {
		return nil
	}
	return &applyment
}

func ApplymentRepositoryFactory() ApplymentRepository {
	repository := ApplymentRepository{
		connection: configs.Connection,
	}
	return repository
}
