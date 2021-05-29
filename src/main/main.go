package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//CONST VARIABLES

//Handler methods
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	timeNowResponse := GetTimeResponse{time.Now().String()}
	b, err := json.Marshal(timeNowResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

//MAIN
func main() {
	r := mux.NewRouter()

	apiv1 := r.PathPrefix("/api/v1").Subrouter()

	apiv1.HandleFunc("/", get).Methods(http.MethodGet)
	apiv1.HandleFunc("/", post).Methods(http.MethodPost)
	apiv1.HandleFunc("/", put).Methods(http.MethodPut)
	apiv1.HandleFunc("/", delete).Methods(http.MethodDelete)
	apiv1.HandleFunc("/", notFound)

	apiv1.HandleFunc("/getTime", getTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":5990", r))
}
