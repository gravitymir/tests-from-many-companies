package httphandlers

import (
	"fmt"
	"net/http"
	"sync"
)

var mtx sync.Mutex

func convertMethod(m string) string {
	switch m {
	case http.MethodPut:
		return "UPDATE"
	case http.MethodDelete:
		return "DELETE"
	case http.MethodPost:
		return "INSERT"
	default:
		return "SELECT"
	}
}

//Job is HTTPHandler for Job
func Job() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mtx.Lock()
		defer mtx.Unlock()

		fmt.Fprintf(w, "Method: "+r.Method+"\n"+convertMethod(r.Method)+" Job")
		fmt.Println("Method: " + r.Method + " " + convertMethod(r.Method) + " Job")
	})
}

//Task is HTTPHandler for Task
func Task() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mtx.Lock()
		defer mtx.Unlock()
		fmt.Fprintf(w, "Method: "+r.Method+"\n"+convertMethod(r.Method)+" Task")
		fmt.Println("Method: " + r.Method + " " + convertMethod(r.Method) + " Task")
	})
}
