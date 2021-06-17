package controller

import (
	"encoding/json"
	"log"
	"moneygement-web/model"
	"moneygement-web/service"
	"net/http"
)

func SavingHandler(w http.ResponseWriter, r *http.Request) {

	// request switcher
	switch r.Method {
	case http.MethodGet:
		savings := findSavingList()
		err := json.NewEncoder(w).Encode(savings)
		if err != nil {
			log.Println("Error: saving value is ", err)
		}
	case http.MethodPost:
		findSaving()
	case http.MethodPatch:
		updateSaving()
	case http.MethodDelete:
		deleteSaving()
	}
}

func findSavingList() *model.SavingList{
	return service.FindSavingList()
}

func findSaving() {
}

func updateSaving() {
}

func deleteSaving() {
}
