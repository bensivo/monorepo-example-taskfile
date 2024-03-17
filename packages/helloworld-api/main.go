package main

import (
	"encoding/json"
	"fmt"
	greetingsv1 "helloworld-proto-go/proto/greetings/v1"
	"io"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	mux := http.NewServeMux()

	// Return CORS for all options requests
	mux.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Options")
		LogRequest(r)
		Cors(w, r)

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /greet", func(w http.ResponseWriter, r *http.Request) {
		LogRequest(r)
		Cors(w, r)

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			WriteJson(w, map[string]any{
				"error": fmt.Errorf("failed to read request body: %w", err),
			})
		}

		// Parse GreetRequest from http request body
		var greetRequest greetingsv1.GreetRequest
		err = proto.Unmarshal(bytes, &greetRequest)
		if err != nil {
			WriteJson(w, map[string]any{
				"error": fmt.Errorf("failed to unmarshal GreetRequest: %w", err),
			})
		}

		// Create GreetResponse
		greetResponse := greetingsv1.GreetResponse{
			Greeting: fmt.Sprintf("Hello, %s!!\n", greetRequest.Name),
		}

		// Write GreetResponse to http response body
		bytes, err = proto.Marshal(&greetResponse)
		if err != nil {
			WriteJson(w, map[string]any{
				"error": fmt.Errorf("failed to marshal GreetResponse: %w", err),
			})
		}
		w.Write(bytes)
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
