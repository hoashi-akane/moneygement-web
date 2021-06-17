package controller

import "net/http"

func UserHandler(w http.ResponseWriter, r *http.Request) {

	// request switcher
	switch r.Method {
	case http.MethodGet:
		findUserList()
	case http.MethodPost:
		findUser()
	case http.MethodPatch:
		updateUser()
	case http.MethodDelete:
		deleteUser()
	}
}

func findUserList() {
}

func findUser() {
}

func updateUser() {
}

func deleteUser() {
}