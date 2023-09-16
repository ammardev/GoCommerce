package http

import "github.com/labstack/echo/v4"

type Context struct {
    echo.Context
}

func (c *Context) Bind(i interface{}) error {
    err := (&echo.DefaultBinder{}).Bind(i, c)
    if err != nil {
        return err
    }

    err = (&echo.DefaultBinder{}).BindHeaders(c, i)
    if err != nil {
        return err
    }

    return nil
}

func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cc := &Context{Context: c}
        return next(cc)
    }
}
