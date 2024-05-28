package applyment_domain

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	applyment_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/applyment"
)

type ListApplyments struct {
	ApplymentRepository applyment_repository.ApplymentRepositoryInterface
}

func (c ListApplyments) Handle() (applyments []*entity.Applyment, errStatus int, err error) {
	applyments, err = c.ApplymentRepository.ListAll()
	if err != nil {
		errStatus = 400
	}
	return
}

func ListApplymentsFactory() ListApplyments {
	return ListApplyments{
		ApplymentRepository: applyment_repository.ApplymentRepositoryFactory(),
	}
}
