package validation

import (
	"encoding/json"
	"errors"
	"github.com/PrrP17/crud-go.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translatin "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translatin.RegisterDefaultTranslations(val, transl)
	}

}

func ValidationError(validation_err error) *rest_err.RestErr {
	var jasonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validation_err, &jasonErr) {
		return rest_err.NewBadRequestErr("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationErr) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationErr("Some field ar invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestErr("Error trying to convert fields")
	}

}
