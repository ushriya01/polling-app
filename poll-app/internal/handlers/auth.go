package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"poll-app/internal/models"
	"poll-app/internal/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// SignupHandler handles user sign up requests
func SignupHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := models.CreateUser(r.Context(), user.Username, string(hashedPassword)); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
		log.Println("Error creating user:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	successResponse := struct {
		Message string `json:"message"`
	}{
		Message: "Signup successful",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successResponse)
}

// SigninHandler handles user sign in requests
func SigninHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}
	userFromDB, err := models.GetUserByUsername(r.Context(), user.Username)
	if err != nil {
		log.Println("Error retrieving user:", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		log.Println("Error comparing passwords:", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(utils.SigningKey)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+tokenString)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
