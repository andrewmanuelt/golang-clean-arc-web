package repository

import "golang-clean-arc/entitiy"

type ExampleRepository interface {
	Get(app_ver string) *entitiy.App
}
