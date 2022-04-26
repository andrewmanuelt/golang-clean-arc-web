package repository

import "gorm.io/gorm"

type IndexRepository interface {
	Index()
}

type indexRepositoryImpl struct {
	Database *gorm.DB
}

func (repository *indexRepositoryImpl) Index() {
}

func NewIndexRepository(database *gorm.DB) IndexRepository {
	return &indexRepositoryImpl{
		Database: database,
	}
}
