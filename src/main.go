package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

type person struct {
	name [][]string
	time [][]string
}

type release struct {
	name   [][]string
	detail [][]string
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/roster", roster).Methods("GET")
	router.HandleFunc("/releases", releases).Methods("GET")
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":4000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I have provided you with the latest data")
	latency := (rand.Float64() * 80) + 80
	traffic := (rand.Float64() * 700) + 50
	errors := (rand.Float64() * 5) + 1
	saturation := (rand.Float64() * 70) + 15
	json.NewEncoder(w).Encode(map[string]float64{"latency": latency, "traffic": traffic, "errors": errors, "saturation": saturation})
}

func releases(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I have provided you with the release data")
	var released = map[string]*release{}
	released["first release"] = &release{name: [][]string{{"Brumbies"}}, detail: [][]string{{"Payroll database upgrade"}}}
	released["second release"] = &release{name: [][]string{{"Identity"}}, detail: [][]string{{"SSO to 20% of all users"}}}
	released["third release"] = &release{name: [][]string{{"Fringe"}}, detail: [][]string{{"Fix deployed for global search"}}}
	json.NewEncoder(w).Encode(map[string][][]string{"team1": released["first release"].name, "team2": released["second release"].name, "team3": released["third release"].name, "detail1": released["first release"].detail, "detail2": released["second release"].detail, "detail3": released["third release"].detail})
}

func roster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I have provided you with the roster data")
	var people = map[string]*person{}
	people["first person"] = &person{name: [][]string{{"A great Xero"}}, time: [][]string{{"10am NZT - 3pm NZT"}}}
	people["second person"] = &person{name: [][]string{{"A fellow Xero"}}, time: [][]string{{"9am NZT - 9pm NZT"}}}
	json.NewEncoder(w).Encode(map[string][][]string{"roster": people["first person"].name, "roster2": people["second person"].name, "time2": people["second person"].time, "time1": people["first person"].time})
}
