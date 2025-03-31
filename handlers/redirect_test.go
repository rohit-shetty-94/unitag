package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRedirectHandler_ValidId(t *testing.T) {
	// Initialize Echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Manually set the path parameter (important step)
	c.SetParamNames("id")
	c.SetParamValues("2bXk5q")

	// Set headers
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("User-Agent", "Chrome")

	// Call the handler
	err := RedirectHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, "https://www.unitag.io/en", rec.Header().Get("Location"))
}

func TestRedirectHandler_InvalidId(t *testing.T) {
	// Initialize Echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set headers
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("User-Agent", "Chrome")

	// Call the handler
	_ = RedirectHandler(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
