package handler

import "net/http"

type commonHeaderHandler struct {
	Next http.Handler
}

func GetCommonHeader(next http.Handler) commonHeaderHandler {
	return commonHeaderHandler{
		Next: next,
	}
}

func (h commonHeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	h.Next.ServeHTTP(w, r)
}
