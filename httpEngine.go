package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func startHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", startCPUstress).Methods("Get")
	r.HandleFunc("/healthcheck", healthCheck).Methods("Get")

	http.ListenAndServe(fmt.Sprintf("%s:%d", *httpAddress, *httpPort), r)
}

func startCPUstress(w http.ResponseWriter, r *http.Request) {

	if hostname, err := os.Hostname(); err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Starting CPU stress on: %s\n", hostname)
		if envData, ok := os.LookupEnv("SPECIAL_DATA"); ok {
			fmt.Fprintf(w, "Specail data is set, value is: %s\n", envData)
		}

		go cpuStress()
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
