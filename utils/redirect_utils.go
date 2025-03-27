package utils

import (
	"strings"
	"unitag/constants"
)

// GetRedirectURL - Determines the redirection URL based on the ID, language, and user-agent
func GetRedirectURL(id, lang, userAgent string) string {
	defaultURL := constants.DefaultURL + lang

	// Identify based on ID and user-agent
	if url, exists := constants.IdentifierMap[id]; exists {
		defaultURL = url
	}

	// User-Agent-based redirection
	if strings.Contains(strings.ToLower(userAgent), "chrome") || strings.Contains(strings.ToLower(userAgent), "android") {
		defaultURL = constants.IdentifierMap[id] + lang
	} else if strings.Contains(strings.ToLower(userAgent), "safari") || strings.Contains(strings.ToLower(userAgent), "iphone") || strings.Contains(strings.ToLower(userAgent), "ipad") {
		defaultURL = constants.IdentifierMap[id]
		if lang != constants.DefaultLanguage {
			defaultURL += lang
		}
	}

	return defaultURL
}
