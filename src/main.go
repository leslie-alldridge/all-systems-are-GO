package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	router.HandleFunc("/message", handleQryMessage).Methods("GET")
	router.HandleFunc("/m/{msg}", handleUrlMessage).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleUrlMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]
	if message == "latency" {
		response := Data{low: 10, high: 10, healthy: 10}
		fmt.Println(response)
		json.NewEncoder(w).Encode(map[string]float64{"low": response.low, "healthy": response.healthy, "high": response.high})

	}
	if message == "traffic" {
		response := Data{low: 650, high: 700, healthy: 650}
		fmt.Println(response)
		json.NewEncoder(w).Encode(map[string]float64{"low": response.low, "healthy": response.healthy, "high": response.high})

	}
	if message == "errors" {
		response := Data{low: 0.9, high: 1.9, healthy: 1.2}
		fmt.Println(response)
		json.NewEncoder(w).Encode(map[string]float64{"low": response.low, "healthy": response.healthy, "high": response.high})

	}
	if message == "saturation" {
		response := Data{low: 10, high: 50, healthy: 35}
		fmt.Println(response)
		json.NewEncoder(w).Encode(map[string]float64{"low": response.low, "healthy": response.healthy, "high": response.high})

	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
