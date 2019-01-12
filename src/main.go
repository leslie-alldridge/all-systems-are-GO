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

type Data struct {
	healthy float64
	low     float64
	high    float64
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")

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
