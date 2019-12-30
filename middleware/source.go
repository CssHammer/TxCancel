package middleware

import (
	"net/http"
)

const HeaderSourceType = "Source-Type"

// SourceType validates source type header
func (m *Middlewares) SourceType(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		sourceType := r.Header.Get(HeaderSourceType)
		if _, ok := m.sourceTypeList[sourceType]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}
