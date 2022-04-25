package repository

import (
	"golang-clean-arc-web/entity"

	"github.com/stretchr/testify/mock"
)

type ExampleRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ExampleRepositoryMock) Get(app_ver string) *entity.App {
	args := repository.Mock.Called(app_ver)

	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entity.App)

	return &category
}
