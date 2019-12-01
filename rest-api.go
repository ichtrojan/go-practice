package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Application Object
type Application struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Developer   string   `json:"developer,omitempty"`
	Details     *Details `json:"details,omitempty"`
}

// Address Object
type Details struct {
	Language string `json:"language,omitempty"`
	Type     string `json:"type,omitempty"`
}

// Define People
var app []Application

// Define port to run
const (
	PORT = ":1909"
)

// GetAppEndPoint A SINGLE APP USING ID
func GetAppEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range app {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Application{})
}

// GetAppsEndpoint GET ALL APPS
func GetAppsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(app)
}

// CreateAppEndpoint CREATE A APP
func CreateAppEndpoint(w http.ResponseWriter, req *http.Request) {
	var appnew Application
	_ = json.NewDecoder(req.Body).Decode(&appnew)
	app = append(app, appnew)
	json.NewEncoder(w).Encode(app)
}

// DeleteAppEndpoint DELETE AN APP
func DeleteAppEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range app {
		if item.ID == params["id"] {
			app = append(app[:index], app[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(app)
}

// ROUTING
func serveRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/app", GetAppsEndpoint).Methods("GET")
	router.HandleFunc("/app/{id}", GetAppEndpoint).Methods("GET")
	router.HandleFunc("/app", CreateAppEndpoint).Methods("POST")
	router.HandleFunc("/app/{id}", DeleteAppEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(PORT, router))

}

func main() {
	app = append(app, Application{ID: "1", Name: "go-api", Description: "Playing around with GO CRUD operations", Developer: "ozombo", Details: &Details{Language: "GO", Type: "Backend/API"}})
	app = append(app, Application{ID: "2", Name: "go-cli", Description: "Cli operations in GO"})
	log.Println("Sample apps populated, your api is serving at: http://127.0.0.1", PORT, "/")
	serveRequests()
}
