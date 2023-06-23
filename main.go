package main

import (
	"log"
	"net/http"
	"os"
	/* These are to be imported when implementing the API

	"github.com/jotacamou/buster/busterapi"
	"github.com/jotacamou/buster/database"

	*/)

var BUSTER_API_URL string

func main() {
	// The minimum requirement is to set the BUSTER_API_URL
	// environment variable, without it all of this is useless.
	if val, ok := os.LookupEnv("BUSTER_API_URL"); !ok {
		log.Fatalf("BUSTER_API_URL not found on runtime")
	} else {
		BUSTER_API_URL = val
	}

	mux := http.NewServeMux()

	// These are the endpoints to be implemented
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/transactions", transactionsHandler)
	mux.HandleFunc("/webhooks", webhookHandler)

	log.Fatal(http.ListenAndServe(":8081", mux))

}

// rootHandler handles the "/" route
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ok": true}`))
}

// transactionsHandler handles the "/transactions" route and only processes
// GET and POST requests.  Anything else results in MethodNotAllowed.
func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok": true}`))
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// webhookHandler handles the "/webhooks" route and only accepts POST
// requests.  Anything else results in InternalServerError.
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, `{"error": "invalid request method"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Handle webhooks
	w.Write([]byte(`{"ok": true}`))
}
