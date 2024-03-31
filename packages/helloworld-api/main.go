package main

import (
	"encoding/json"
	"fmt"
	"helloworld-core/pkg/greetings"
	"net/http"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Return CORS for all options requests
	mux.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		LogRequest(r)
		Cors(w, r)

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /greetings/{name}", func(w http.ResponseWriter, r *http.Request) {
		LogRequest(r)
		Cors(w, r)

		name := r.PathValue("name")
		WriteJson(w, map[string]any{
			// "greeting": greetings.Greet(name),
			"greeting": strings.ToUpper(greetings.Greet(name)),
		})
	})

	fmt.Println("Starting HTTP server on localhost:3000")
	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		fmt.Printf("Error starting http server: %v\n", err)
	}
}

func LogRequest(r *http.Request) {
	fmt.Printf("%s %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
}

func WriteJson(w http.ResponseWriter, body map[string]any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{
			"error":"JSON marshall failed",	
		}`))
	}

	w.Write(bytes)
}

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow from any origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow specific headers
}
