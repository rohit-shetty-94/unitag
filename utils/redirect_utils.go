package utils

import "strings"

// GetRedirectURL - Determines the redirection URL based on the ID, language, and user-agent
func GetRedirectURL(id, lang, userAgent string) string {
	defaultURL := "https://www.unitag.io/" + lang //const

	IdentifierMap := map[string]string{
		"a2Rty1": "https://www.unitag.io/",
		"ax3Goo": "https://www.google.com/?hl=",
		"ax4App": "https://www.apple.com/",
	}

	// Identify based on ID and user-agent
	if url, exists := IdentifierMap[id]; exists {
		defaultURL = url
	}

	// User-Agent-based redirection
	if strings.Contains(strings.ToLower(userAgent), "chrome") || strings.Contains(strings.ToLower(userAgent), "android") {
		defaultURL = IdentifierMap[id] + lang
	} else if strings.Contains(strings.ToLower(userAgent), "safari") || strings.Contains(strings.ToLower(userAgent), "iphone") || strings.Contains(strings.ToLower(userAgent), "ipad") {
		defaultURL = IdentifierMap[id]
		if lang != "en" { //const
			defaultURL += lang
		}
	}

	return defaultURL
}
