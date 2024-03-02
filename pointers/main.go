package main

import (
	"encoding/json"
	"go-playground/pointers/cmd/api"
	"log"
	"net/http"
	"time"
)

type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		hc := HealthCheck{
			Status:    "ok",
			Timestamp: time.Now().GoString(),
		}
		response, _ := json.Marshal(hc)
		w.Write(response)
	})
	http.HandleFunc("/items/create", api.HandleCreate)
	http.HandleFunc("/items", api.HandleList)
	http.HandleFunc("/items/details", api.HandleRead)
	http.HandleFunc("/items/update", api.HandleUpdate)
	http.HandleFunc("/items/delete", api.HandleDelete)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
