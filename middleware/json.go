package middleware

import (
	"net/http"
)

const (
	HeaderContentType = "Content-Type"
	ContentTypeJson   = "application/json"
)

// JSON sets content type of response
func (m *Middlewares) JSON(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentType, ContentTypeJson)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}
