package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"github.com/TMg00000/customerscheduleapi/internal/domain/models/resources/resourceserrorsmessages"
	"github.com/TMg00000/customerscheduleapi/internal/validation"
)

var DataBase []requests.Client

func CreateAppointments(w http.ResponseWriter, r *http.Request) {
	var client requests.Client

	client.Id = requests.NewId()

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, resourceserrorsmessages.BadRequest, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validationErrors := validation.ListErrorsMessages(client)
	if len(validationErrors) > 0 {
		msg, _ := json.MarshalIndent(validationErrors, "", "  ")
		http.Error(w, string(msg), http.StatusBadRequest)

		return
	}

	DataBase = append(DataBase, client)

	db, _ := json.MarshalIndent(DataBase, "", "  ")
	fmt.Printf("%s", string(db))

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(DataBase)
}

func CatchAllAppointments(w http.ResponseWriter, r *http.Request) {
	db, _ := json.MarshalIndent(DataBase, "", "  ")
	fmt.Printf("%s", string(db))

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(DataBase)
}

func UpdateAppointmentsByid(w http.ResponseWriter, r *http.Request) {
	var updateClient requests.UpdateClient

	if err := json.NewDecoder(r.Body).Decode(&updateClient); err != nil {
		http.Error(w, resourceserrorsmessages.BadRequest, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validationErrors := validation.ListErrorsMessages(updateClient)
	if len(validationErrors) > 0 {
		msg, _ := json.MarshalIndent(validationErrors, "", "  ")
		http.Error(w, string(msg), http.StatusBadRequest)

		return
	}

	found := false
	for i, c := range DataBase {
		if c.Id == updateClient {
			DataBase[i] = updateClient
			found = true
			break
		}
	}

	if !found {
		http.Error(w, resourceserrorsmessages.NotFound, http.StatusNotFound)
	}

	db, _ := json.MarshalIndent(DataBase, "", "  ")
	fmt.Printf("%s", string(db))

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(DataBase)
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	var del requests.DeleteAppointment

	found := false
	for i, c := range DataBase {
		if c.Id == del {
			DataBase = append(DataBase[:i], DataBase[i+1])
			found = true
			break
		}
	}

	if !found {
		http.Error(w, resourceserrorsmessages.NotFound, http.StatusNotFound)
	}

	db, _ := json.MarshalIndent(DataBase, "", "  ")
	fmt.Printf("%s", string(db))

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(DataBase)
}
