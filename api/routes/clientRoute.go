package routes

import (
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"formacao-go-dio/api/connections"
	"formacao-go-dio/api/models"
)

func ClientRoute(router *mux.Router, basePath string) *mux.Router {
	router.HandleFunc(basePath+"/clients", GetClientsHandler).Methods("GET")
	router.HandleFunc(basePath+"/clients/{id}", GetClientByIDHandler).Methods("GET")
	router.HandleFunc(basePath+"/clients", CreateClientHandler).Methods("POST")
	router.HandleFunc(basePath+"/clients/{id}", UpdateClientHandler).Methods("PUT")
	router.HandleFunc(basePath+"/clients/{id}", DeleteClientHandler).Methods("DELETE")

	return router
}

func GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connections.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Getting clients on database...")
	clients := models.GetClients(db)

	if clients == nil {
		http.Error(w, "No clients found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func GetClientByIDHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connections.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Getting client by id on database...")
	vars := mux.Vars(r)
	id := vars["id"]
	clientID, err := strconv.Atoi(id)
	client, err := models.GetClientByID(db, int64(clientID))
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connections.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Creating client on database...")
	var client models.Client
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := models.CreateClient(db, client)
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connections.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Updating client on database...")
	vars := mux.Vars(r)
	id := vars["id"]
	clientID, err := strconv.Atoi(id)
	client, err := models.GetClientByID(db, int64(clientID))
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.UpdateClient(db, client)
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connections.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Deleting client on database...")
	vars := mux.Vars(r)
	id := vars["id"]
	clientID, err := strconv.Atoi(id)
	client, err := models.GetClientByID(db, int64(clientID))
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.DeleteClient(db, client.ID)
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}