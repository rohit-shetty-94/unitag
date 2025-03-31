package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for GetRedirectURL function
func TestGetRedirectURL(t *testing.T) {
	tests := []struct {
		id          string
		lang        string
		userAgent   string
		expectedURL string
	}{
		{"a2Rty1", "en", "Chrome", "https://www.unitag.io/en"},
		{"ax3Goo", "fr", "Safari", "https://www.google.com/?hl=fr"},
		{"ax4App", "es", "Ipad", "https://www.apple.com/es"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			redirectURL := GetRedirectURL(tt.id, tt.lang, tt.userAgent)
			assert.Equal(t, tt.expectedURL, redirectURL)
		})
	}
}
