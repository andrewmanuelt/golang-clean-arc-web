package service

import (
	"errors"
	"golang-clean-arc-web/entitiy"
	"golang-clean-arc-web/test/repository"
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
