package auth

import (
	"context"
	"encoding/json"
	"hls_zappingtv/db"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	_, err := db.DB.Exec(context.Background(), `INSERT INTO users (username, password) VALUES ($1, $2)`, u.Username, u.Password)

	if err != nil {
		http.Error(w, "Error al registrar usuario", http.StatusBadRequest)
		log.Printf("Error al registrar usuario: %v", err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	var dbPass string
	err := db.DB.QueryRow(context.Background(),
		`SELECT password FROM users WHERE username=$1`, u.Username).Scan(&dbPass)

	if err != nil || dbPass != u.Password {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Login exitoso"})
}
