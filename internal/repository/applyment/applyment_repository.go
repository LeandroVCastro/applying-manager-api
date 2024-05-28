package applyment_repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type ApplymentRepositoryInterface interface {
	GetById(id uint) *entity.Applyment
	CreateOrUpdate(id uint, title string, description, link *string, company_id, platform_id *uint, applied_at *time.Time) (uint, error)
	ListAll() (applyments []*entity.Applyment, err error)
	Delete(id uint) error
}

type ApplymentRepository struct {
	connection *gorm.DB
}

func (repository ApplymentRepository) CreateOrUpdate(
	id uint,
	title string,
	description *string,
	link *string,
	company_id *uint,
	platform_id *uint,
	applied_at *time.Time,
) (savedId uint, err error) {
	applymentParams := entity.Applyment{Title: title}
	if *description != "" {
		applymentParams.Description = description
	} else {
		applymentParams.Description = nil
	}

	if *link != "" {
		applymentParams.Link = link
	} else {
		applymentParams.Link = nil
	}

	if *company_id != 0 {
		applymentParams.CompanyId = company_id
	} else {
		applymentParams.CompanyId = nil
	}

	if *platform_id != 0 {
		applymentParams.PlatformId = platform_id
	} else {
		applymentParams.PlatformId = nil
	}

	if !time.Time.IsZero(*applied_at) {
		applymentParams.AppliedAt = applied_at
	} else {
		applymentParams.AppliedAt = nil
	}

	var result *gorm.DB
	if id != 0 {
		applymentParams.ID = id
		result = repository.connection.Model(applymentParams).
			Update("title", applymentParams.Title).
			Update("description", applymentParams.Description).
			Update("link", applymentParams.Link).
			Update("company_id", applymentParams.CompanyId).
			Update("platform_id", applymentParams.PlatformId).
			Update("applied_at", applymentParams.AppliedAt)
	} else {
		result = repository.connection.Create(&applymentParams)
	}
	if result.Error != nil {
		fmt.Println("Error on createOrUpdate ApplymentRepository: " + result.Error.Error())
		err = result.Error
		return
	}
	savedId = applymentParams.ID
	return
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
	return ApplymentRepository{
		connection: configs.Connection,
	}
}
