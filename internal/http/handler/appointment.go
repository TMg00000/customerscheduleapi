package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/requests"
	"github.com/TMg00000/customerscheduleapi/internal/domain/models/resources/resourceserrorsmessages"
	"github.com/TMg00000/customerscheduleapi/internal/repository"
	"github.com/TMg00000/customerscheduleapi/internal/validation"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAppointments(repo *repository.AppointmentsRepository, counters *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client requests.Client

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

		if err := repo.Create(client); err != nil {
			http.Error(w, resourceserrorsmessages.ErrorAddInDataBase, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(client)
	}
}

func GetAllAppointments(repo *repository.AppointmentsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := repo.GetAll()
		if err != nil {
			http.Error(w, resourceserrorsmessages.ErroQueryDataBase, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(clients)
	}
}

func UpdateAppointments(repo *repository.AppointmentsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateClient requests.Client
		if err := json.NewDecoder(r.Body).Decode(&updateClient); err != nil {
			http.Error(w, resourceserrorsmessages.NotFound, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		vars := mux.Vars(r)

		idStr := vars["id"]
		id, errparts := primitive.ObjectIDFromHex(idStr)
		if errparts != nil {
			http.Error(w, resourceserrorsmessages.IdInvalid, http.StatusNotFound)
			return
		}
		updateClient.Id = id

		validationErrors := validation.ListErrorsMessages(updateClient)
		if len(validationErrors) > 0 {
			msg, _ := json.MarshalIndent(validationErrors, "", "  ")
			http.Error(w, string(msg), http.StatusBadRequest)
			return
		}

		if err := repo.Update(updateClient); err != nil {
			http.Error(w, resourceserrorsmessages.ErrorUpdateInDataBase, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updateClient)
	}
}

func DeleteAppointments(repo *repository.AppointmentsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deleteClient requests.Client
		if err := json.NewDecoder(r.Body).Decode(&deleteClient); err != nil {
			http.Error(w, resourceserrorsmessages.NotFound, http.StatusNotFound)
			return
		}
		defer r.Body.Close()

		vars := mux.Vars(r)

		idStr := vars["id"]
		id, errparts := primitive.ObjectIDFromHex(idStr)
		if errparts != nil {
			http.Error(w, resourceserrorsmessages.IdInvalid, http.StatusNotFound)
			return
		}

		if err := repo.Delete(id); err != nil {
			http.Error(w, resourceserrorsmessages.ErrorUpdateInDataBase, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deleteClient)
	}
}
