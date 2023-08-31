package server

import (
	"database/sql"
	net_http "net/http"

	"github.com/ammardev/gocommerce/internal/http"
	"github.com/labstack/echo/v4"
)

func httpErrorHandler(err error, c echo.Context) {
	err = handleByErrorType(err, c)
	if err == nil {
		return
	}

	err = handleByErrorValue(err, c)
	if err == nil {
		return
	}

	c.JSON(echo.ErrInternalServerError.Code, echo.ErrInternalServerError)
	c.Logger().Error(err)
}

func handleByErrorType(err error, c echo.Context) error {
	switch err := err.(type) {
	case *http.ValidationErrors:
		c.JSON(net_http.StatusUnprocessableEntity, err)
	case *echo.HTTPError:
		c.JSON(err.Code, echo.HTTPError{
			Message: err.Error(),
		})
	default:
		// Return the error again if it wasn't handled
		return err
	}

	return nil
}

func handleByErrorValue(err error, c echo.Context) error {
	switch err {
	case sql.ErrNoRows:
		c.JSON(echo.ErrNotFound.Code, echo.ErrNotFound)
	default:
		// Return the error again if it wasn't handled
		return err
	}

	return nil
}
