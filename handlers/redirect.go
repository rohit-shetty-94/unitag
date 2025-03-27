package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func RedirectHandler(c echo.Context) error {

	redirectURL := ""
	return c.Redirect(http.StatusFound, redirectURL)

}
