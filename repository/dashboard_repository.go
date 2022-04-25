package repository

import "gorm.io/gorm"

type DashboardRepository interface {
	Home()
}

type dashboardRepositoryImpl struct {
	Database *gorm.DB
}

func (repository *dashboardRepositoryImpl) Home() {

}

func NewDashboardRepository(database *gorm.DB) DashboardRepository {
	return &dashboardRepositoryImpl{
		Database: database,
	}
}
