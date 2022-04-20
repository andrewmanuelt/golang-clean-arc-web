package helper

import (
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validates(data interface{}) interface{} {
	id := id.New()

	uni = ut.New(id, id)

	trans, _ := uni.GetTranslator("id")

	validate = validator.New()

	id_translations.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(data); err != nil {

		errs := err.(validator.ValidationErrors)

		tag := errs.Translate(trans)

		return tag
	}

	return nil
}
