package server

import (
	"net/http"

	"github.com/ammardev/gocommerce/app"
	"github.com/labstack/echo/v4"
)

func httpErrorHandler(err error, c echo.Context) {
	skipLogging := false

	switch err.(type) {
	case *app.ValidationErrors:
		c.JSON(http.StatusUnprocessableEntity, err)
		skipLogging = true
	case *echo.HTTPError:
		c.JSON(http.StatusInternalServerError, echo.HTTPError{
			Message: err.Error(),
		})
		skipLogging = true
	default:
		c.JSON(http.StatusInternalServerError, echo.HTTPError{
			Message: "Internal Error",
		})
	}

	if !skipLogging {
		c.Logger().Error(err)
	}
}