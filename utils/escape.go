package utils

import (
	"net/url"
	"strings"
)

func URLQueryEscape(origin string) string {
	escaped := url.QueryEscape(origin)
	escaped = strings.Replace(escaped, "%2F", "/", -1)
	escaped = strings.Replace(escaped, "%3D", "=", -1)
	escaped = strings.Replace(escaped, "+", "%20", -1)
	return escaped
}

func URLQueryUnescape(escaped string) (string, error) {
	escaped = strings.Replace(escaped, "/", "%2F", -1)
	escaped = strings.Replace(escaped, "=", "%3D", -1)
	escaped = strings.Replace(escaped, "%20", " ", -1)
	return url.QueryUnescape(escaped)
}
