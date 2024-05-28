package applyment_domain

import (
	"errors"
	"time"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/applyment"
)

type SaveApplyment struct {
	ApplymentRepository applyment_repository.ApplymentRepositoryInterface
}

func (a SaveApplyment) Handle(
	id uint,
	title string,
	description *string,
	link *string,
	company_id *uint,
	platform_id *uint,
	applied_at *time.Time,
) (savedCompany *entity.Applyment, errStatus int, err error) {
	if id != 0 {
		if company := a.ApplymentRepository.GetById(id); company == nil {
			err = errors.New("applyment not found")
			errStatus = 404
			return
		}
		savedId, updatedErr := a.ApplymentRepository.CreateOrUpdate(id, title, description, link, company_id, platform_id, applied_at)
		if updatedErr != nil {
			err = errors.New("error on update applyment")
			errStatus = 400
			return
		}
		savedCompany = a.ApplymentRepository.GetById(savedId)
		return
	}
	savedId, saveErr := a.ApplymentRepository.CreateOrUpdate(0, title, description, link, company_id, platform_id, applied_at)
	if saveErr != nil {
		err = errors.New("error on create applyment")
		errStatus = 400
		return
	}
	savedCompany = a.ApplymentRepository.GetById(savedId)
	return
}

func SaveApplymentFactory() SaveApplyment {
	return SaveApplyment{
		ApplymentRepository: applyment_repository.ApplymentRepositoryFactory(),
	}
}
