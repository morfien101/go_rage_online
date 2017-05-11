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

	http.ListenAndServe(fmt.Sprintf("%s:%d", *httpAddress, *httpPort), r)
}

func startCPUstress(w http.ResponseWriter, r *http.Request) {

	if hostname, err := os.Hostname(); err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Starting CPU stress on: %s\n", hostname)
		go cpuStress()
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
