package service

import "golang-clean-arc-web/repository"

type IndexService interface {
	Index()
}

type indexServiceImpl struct {
	IndexRepository repository.IndexRepository
}

func (service *indexServiceImpl) Index() {}

func NewIndexService(indexRepository *repository.IndexRepository) IndexService {
	return &indexServiceImpl{
		IndexRepository: *indexRepository,
	}
}
