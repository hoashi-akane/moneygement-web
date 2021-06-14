package controller

import "net/http"

func UserHandler(w http.ResponseWriter, r *http.Request) {

	// request switcher
	switch r.Method {
	case http.MethodGet:
		findAll(w, r)
	case http.MethodPost:

	case http.MethodPatch:

	case http.MethodDelete:

	}
}

func findAll(w http.ResponseWriter, r *http.Request) {

}