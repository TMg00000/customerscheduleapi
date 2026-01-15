package main

import (
	"net/http"

	"github.com/TMg00000/customerscheduleapi/internal/http/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/toschedule", handler.CreateAppointments).Methods("POST")
	r.HandleFunc("/catchallappointments", handler.CatchAllAppointments).Methods("GET")
	http.ListenAndServe(":9437", r)
}
