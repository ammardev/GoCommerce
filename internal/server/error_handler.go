package server

import (
	net_http "net/http"

	"github.com/ammardev/gocommerce/internal/http"
	"github.com/labstack/echo/v4"
)

func httpErrorHandler(err error, c echo.Context) {
	skipLogging := false

	switch err.(type) {
	case *http.ValidationErrors:
		c.JSON(net_http.StatusUnprocessableEntity, err)
		skipLogging = true
	case *echo.HTTPError:
		c.JSON(net_http.StatusInternalServerError, echo.HTTPError{
			Message: err.Error(),
		})
		skipLogging = true
	default:
		c.JSON(net_http.StatusInternalServerError, echo.HTTPError{
			Message: "Internal Error",
		})
	}

	if !skipLogging {
		c.Logger().Error(err)
	}
}
