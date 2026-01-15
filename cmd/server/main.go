package main

import (
	"log"
	"net/http"

	"github.com/TMg00000/customerscheduleapi/internal/domain/models/resources/resourceserrorsmessages"
	"github.com/TMg00000/customerscheduleapi/internal/http/handler"
	"github.com/TMg00000/customerscheduleapi/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	client, err := repository.NewMongoDB()
	if err != nil {
		log.Fatal(resourceserrorsmessages.NotFound)
	}

	repo := repository.NewAppointmentsRepository(client)

	r := mux.NewRouter()
	r.HandleFunc("/appointments", handler.CreateAppointments(repo)).Methods("POST")
	r.HandleFunc("/appointments", handler.GetAllAppointments(repo)).Methods("GET")
	r.HandleFunc("/appointments/{id}", handler.UpdateAppointments(repo)).Methods("PUT")
	r.HandleFunc("/appointments/{id}", handler.DeleteAppointments(repo)).Methods("DELETE")
	http.ListenAndServe(":9437", r)
}
