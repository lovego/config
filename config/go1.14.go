// +build go1.14

package config

import "net/http"

func setSameSiteMode(cookie *http.Cookie, sameSite string) {
	sameSiteMode := http.SameSiteDefaultMode
	switch sameSite {
	case "lax":
		sameSiteMode = http.SameSiteLaxMode
	case "strict":
		sameSiteMode = http.SameSiteStrictMode
	case "none":
		sameSiteMode = http.SameSiteNoneMode
	}
	cookie.SameSite = sameSiteMode
}
