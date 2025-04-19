package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/login", loginHandler)

	log.Println("Backend draait op http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Alleen POST toegestaan", http.StatusMethodNotAllowed)
		return
	}

	var creds LoginRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Ongeldige JSON", http.StatusBadRequest)
		return
	}

	if creds.Username == "admin" && creds.Password == "wachtwoord123" {
		json.NewEncoder(w).Encode(LoginResponse{Message: "Login gelukt!"})
	} else {
		http.Error(w, "Ongeldige gebruikersnaam of wachtwoord", http.StatusUnauthorized)
	}
}
