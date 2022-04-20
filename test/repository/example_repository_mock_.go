package repository

import (
	"golang-clean-arc/entitiy"

	"github.com/stretchr/testify/mock"
)

type ExampleRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ExampleRepositoryMock) Get(app_ver string) *entitiy.App {
	args := repository.Mock.Called(app_ver)

	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entitiy.App)

	return &category
}
