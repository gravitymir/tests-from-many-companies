package handlers

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (j *Handler) Job() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: "+r.Method+"\nDelete Job")
	})
}

func (j *Handler) Task() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: "+r.Method+"\nDelete Job")
	})
}

// func (h *Job) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	h.n++
// 	fmt.Fprintf(w, "count is %d\n", h.n)
// }
