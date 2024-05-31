package stage_domain

import (
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	stage_repository "github.com/LeandroVCastro/applying-manager-api/internal/repository/stage"
)

type ListStages struct {
	StageRepository stage_repository.StageRepositoryInterface
}

func (s ListStages) Handle() (stages []*entity.Stage, errStatus int, err error) {
	stages, err = s.StageRepository.ListAll()
	if err != nil {
		errStatus = 400
	}
	return
}

func ListStagesFactory() ListStages {
	return ListStages{
		StageRepository: stage_repository.StageRepositoryFactory(),
	}
}
