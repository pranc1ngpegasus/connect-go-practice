package middleware

import "net/http"

func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	if len(m) == 0 {
		return h
	}

	return m[0](Chain(h, m[1:]...))
}
