package applyment_domain

import (
	"errors"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/applyment"
)

type GetApplyment struct {
	ApplymentRepository applyment_repository.ApplymentRepositoryInterface
}

func (a GetApplyment) Handle(id uint) (applyment *entity.Applyment, errStatus int, err error) {
	if id == 0 {
		err = errors.New("should be provided an ID")
		errStatus = 400
		return
	}
	applyment = a.ApplymentRepository.GetById(id)
	if applyment == nil {
		err = errors.New("applyment not found")
		errStatus = 404
		return
	}
	return

}

func GetApplymentFactory() GetApplyment {
	return GetApplyment{
		ApplymentRepository: applyment_repository.ApplymentRepositoryFactory(),
	}
}
