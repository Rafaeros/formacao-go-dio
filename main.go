package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"formacao-go-dio/Challenges"
	"formacao-go-dio/api/routes"
)

func main() {
	//Challenges.CalcTemperature()
	Challenges.NumberIssue1()
	//Challenges.NumberIssue2()
	//Challenges.PingPong()
	basePath := "/api/v1"

	r := mux.NewRouter()
	log.Println("Setting up routers")
	r = routes.ClientRoute(r, basePath)

	log.Println("Setting up handlers")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/'")
		w.Write([]byte("Hello World"))
	}).Methods(http.MethodGet)

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/ping'")
		w.Write([]byte("pong")) 
	}).Methods(http.MethodGet)

	r.HandleFunc(basePath+"/version", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/version'")
		w.Write([]byte("1.0.0"))
	}).Methods(http.MethodGet)

	log.Printf("Server running on http://localhost:8090%s\n", basePath)
	log.Fatal(http.ListenAndServe(":8090", r))
}