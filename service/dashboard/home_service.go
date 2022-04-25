package dashboard

import "golang-clean-arc-web/repository"

type HomeService interface {
	Home()
}

type homeServiceImpl struct {
	DashboardRepository repository.DashboardRepository
}

func (homeService *homeServiceImpl) Home() {

}

func NewHomeService(dashboardRepository *repository.DashboardRepository) HomeService {
	return &homeServiceImpl{
		DashboardRepository: *dashboardRepository,
	}
}
