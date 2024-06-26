package applyment_domain

import (
	"errors"
	"time"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/applyment"
	company_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/company"
	platform_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/platform"
)

type SaveApplyment struct {
	ApplymentRepository applyment_repository.ApplymentRepositoryInterface
	CompanyRepository   company_repository.CompanyRepositoryInterface
	PlatformRepository  platform_repository.PlatformRepositoryInterface
}

func (a SaveApplyment) Handle(
	id uint,
	title string,
	description *string,
	link *string,
	company_id *uint,
	platform_id *uint,
	applied_at *time.Time,
) (savedApplyment *entity.Applyment, errStatus int, err error) {
	err = a.getCompany(company_id)
	if err != nil {
		errStatus = 404
		return
	}
	err = a.getPlatform(platform_id)
	if err != nil {
		errStatus = 404
		return
	}
	if id != 0 {
		err = a.getApplyment(id)
		if err != nil {
			errStatus = 404
			return
		}
		savedId, updatedErr := a.ApplymentRepository.CreateOrUpdate(id, title, description, link, company_id, platform_id, applied_at)
		if updatedErr != nil {
			err = errors.New("error on update applyment")
			errStatus = 400
			return
		}
		savedApplyment = a.ApplymentRepository.GetById(savedId)
		return
	}
	savedId, saveErr := a.ApplymentRepository.CreateOrUpdate(0, title, description, link, company_id, platform_id, applied_at)
	if saveErr != nil {
		err = errors.New("error on create applyment")
		errStatus = 400
		return
	}
	savedApplyment = a.ApplymentRepository.GetById(savedId)
	return
}

func SaveApplymentFactory() SaveApplyment {
	return SaveApplyment{
		ApplymentRepository: applyment_repository.ApplymentRepositoryFactory(),
		CompanyRepository:   company_repository.CompanyRepositoryFactory(),
		PlatformRepository:  platform_repository.PlatformRepositoryFactory(),
	}
}

func (a SaveApplyment) getApplyment(id uint) error {
	applyment := a.ApplymentRepository.GetById(id)
	if applyment == nil {
		return errors.New("applyment not found")
	}
	return nil
}

func (a SaveApplyment) getCompany(id *uint) error {
	if *id != 0 {
		company := a.CompanyRepository.GetById(*id)
		if company == nil {
			return errors.New("company not found")
		}
	}
	return nil
}

func (a SaveApplyment) getPlatform(id *uint) error {
	if *id != 0 {
		platform := a.PlatformRepository.GetById(*id)
		if platform == nil {
			return errors.New("platform not found")
		}
	}
	return nil
}
