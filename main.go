package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Result      string `json:"result"`
	IsSuperuser bool   `json:"is_superuser"`
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	// Load username and password from environment variables
	username := os.Getenv("EMQX_USERNAME")
	password := os.Getenv("EMQX_PASSWORD")

	// Set content type json
	w.Header().Set("Content-Type", "application/json")

	// Read the request body
	var authRequest AuthRequest
	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Authentication attempt: %+v\n", authRequest)

	// Check if the username and password are correct
	var authResponse AuthResponse
	if authRequest.Username == username && authRequest.Password == password {
		authResponse = AuthResponse{Result: "allow", IsSuperuser: false}
		w.WriteHeader(http.StatusOK)
	} else {
		authResponse = AuthResponse{Result: "deny", IsSuperuser: false}
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(authResponse)
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", authHandler)
	err := http.ListenAndServe(":31337", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
