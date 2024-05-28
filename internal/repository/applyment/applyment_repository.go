package applyment_repository

import (
	"errors"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type ApplymentRepositoryInterface interface {
	GetById(id uint) *entity.Applyment
	ListAll() (applyments []*entity.Applyment, err error)
	Delete(id uint) error
}

type ApplymentRepository struct {
	connection *gorm.DB
}

func (repository ApplymentRepository) ListAll() (applyments []*entity.Applyment, err error) {
	result := repository.connection.Select([]string{"id", "title", "description", "link", "company_id", "platform_id", "applied_at", "created_at", "updated_at"}).Order("id DESC").Find(&applyments)
	if result.Error != nil {
		err = errors.New(result.Error.Error())
		return
	}
	return
}

func (repository ApplymentRepository) GetById(id uint) *entity.Applyment {
	var applyment = entity.Applyment{}
	result := repository.connection.First(&applyment, id)
	if result.Error != nil {
		return nil
	}
	return &applyment
}

func (repository ApplymentRepository) Delete(id uint) error {
	var applyment = entity.Applyment{}
	result := repository.connection.Where("ID = ?", id).Delete(&applyment)
	if result.Error != nil {
		err := errors.New(result.Error.Error())
		return err
	}
	return nil
}

func ApplymentRepositoryFactory() ApplymentRepository {
	repository := ApplymentRepository{
		connection: configs.Connection,
	}
	return repository
}
