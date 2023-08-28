package app

import (
	"github.com/go-playground/locales/ar"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ar_translations "github.com/go-playground/validator/v10/translations/ar"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

var httpValidator *validator.Validate
var translator *ut.UniversalTranslator

type ValidationError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func Validate(c echo.Context, i interface{}) error {
	if err := httpValidator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return newValidationError(&validationErrors, c.Request().Header.Get("Accept-Language"))
	}

	return nil
}

func (err *ValidationError) Error() string {
	return err.Message
}

func newValidationError(validationErrors *validator.ValidationErrors, locale string) *ValidationError {
	errors := make(map[string]string, len(*validationErrors))
	localeTranslator, _ := translator.GetTranslator(locale)

	for _, validationError := range *validationErrors {
		errors[validationError.Field()] = validationError.Translate(localeTranslator)
	}

	return &ValidationError{
		Message: "You have errors in some fields.",
		Errors:  errors,
	}
}

func NewValidator() {
	httpValidator = validator.New()

	translator = ut.New(
		en.New(),
		ar.New(),
	)

	enTranslator, _ := translator.GetTranslator("en")
	arTranslator, _ := translator.GetTranslator("ar")

	en_translations.RegisterDefaultTranslations(httpValidator, enTranslator)
	ar_translations.RegisterDefaultTranslations(httpValidator, arTranslator)
}
