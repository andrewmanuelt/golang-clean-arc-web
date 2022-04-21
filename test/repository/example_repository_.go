package repository

import "golang-clean-arc-web/entitiy"

type ExampleRepository interface {
	Get(app_ver string) *entitiy.App
}
