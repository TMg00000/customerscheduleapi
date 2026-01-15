package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TMg00000/customerscheduleapi/internal/domain/requests"
)

func CreateAppointments(w http.ResponseWriter, r *http.Request) {
	var client requests.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)

}

func ListAppointments(w http.ResponseWriter, r *http.Request) {
	lista := []requests.Client{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}
