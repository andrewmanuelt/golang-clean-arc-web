package service

import (
	"errors"
	"golang-clean-arc/entitiy"
	"golang-clean-arc/test/repository"
)

type ExampleService struct {
	Repository repository.ExampleRepository
}

func (service ExampleService) Get(app_ver string) (*entitiy.App, error) {
	app := service.Repository.Get(app_ver)

	if app == nil {
		return nil, errors.New("app not found")
	}

	return app, nil
}
