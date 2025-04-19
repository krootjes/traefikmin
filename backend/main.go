package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtSecret = []byte("geheim123")

func main() {
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/me", requireAuth(currentUserHandler))

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

	if creds.Username == "daan" && creds.Password == "a" {
		token, err := generateJWT(creds.Username)
		if err != nil {
			http.Error(w, "JWT genereren mislukt", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	} else {
		http.Error(w, "Ongeldige gebruikersnaam of wachtwoord", http.StatusUnauthorized)
	}
}

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || len(auth) < 8 || auth[:7] != "Bearer " {
			http.Error(w, "Geen of ongeldige Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := auth[7:]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			log.Println("JWT fout:", err)
			http.Error(w, "Token ongeldig of verlopen", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		// username doorgeven via context
		ctx := context.WithValue(r.Context(), "username", username)
		next(w, r.WithContext(ctx))
	}
}

func currentUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	json.NewEncoder(w).Encode(map[string]string{
		"username": username,
	})
}
