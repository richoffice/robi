package robi

import (
	"github.com/labstack/echo/v4"
)

type RobiEcho struct {
}

func (re *RobiEcho) New() *echo.Echo {
	e := echo.New()
	return e
}
