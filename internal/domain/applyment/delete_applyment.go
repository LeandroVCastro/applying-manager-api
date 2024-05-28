package applyment_domain

import (
	"errors"

	applyment_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/applyment"
)

type DeleteApplyment struct {
	ApplymentRepository applyment_repository.ApplymentRepositoryInterface
}

func (a DeleteApplyment) Handle(id uint) (errStatus int, err error) {
	if id == 0 {
		err = errors.New("should be provided an ID")
		errStatus = 400
		return
	}
	company := a.ApplymentRepository.GetById(id)

	if company == nil {
		err = errors.New("applyment not found")
		errStatus = 404
		return
	}
	errDelete := a.ApplymentRepository.Delete(id)
	if errDelete != nil {
		err = errors.New(errDelete.Error())
		errStatus = 500
		return
	}
	return
}

func DeleteApplymentFactory() DeleteApplyment {
	return DeleteApplyment{
		ApplymentRepository: applyment_repository.ApplymentRepositoryFactory(),
	}
}
