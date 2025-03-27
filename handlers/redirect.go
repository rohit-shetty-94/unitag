package handlers

import (
	"net/http"
	"strings"
	"unitag/constants"
	"unitag/utils"

	"github.com/labstack/echo"
)

// RedirectHandler - Handles the redirection logic
func RedirectHandler(c echo.Context) error {
	id := c.Param("id")

	// Validate 6-character ID
	if len(id) != 6 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	// Get language & user-agent from headers
	langHeader := c.Request().Header.Get("Accept-Language")
	userAgent := c.Request().Header.Get("User-Agent")

	// Extract primary language (first 2 characters)
	lang := constants.DefaultLanguage
	if len(langHeader) >= 2 {
		lang = strings.ToLower(langHeader[:2])
	}

	// Determine redirection URL using utility function
	redirectURL := utils.GetRedirectURL(id, lang, userAgent)

	// Redirect to the determined URL
	return c.Redirect(http.StatusFound, redirectURL)
}
