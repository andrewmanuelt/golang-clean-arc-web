package service

import (
	"golang-clean-arc/entitiy"
	"golang-clean-arc/test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var exampleRepository = &repository.ExampleRepositoryMock{Mock: mock.Mock{}}
var exampleService = ExampleService{
	Repository: exampleRepository,
}

func TestCategoryService_Get(t *testing.T) {
	exampleRepository.Mock.On("Get", "1.0").Return(nil)

	app, err := exampleService.Get("1.0")

	assert.Nil(t, app)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	app := entitiy.App{
		AppName: "Golang",
		AppVer:  "1",
	}

	exampleRepository.Mock.On("Get", "1").Return(app)

	result, err := exampleService.Get("1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, app.AppName, result.AppName)
}
