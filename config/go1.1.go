// +build go1.1,!go1.12

package config

import "net/http"

func setSameSiteMode(cookie *http.Cookie, sameSite string) {
}
