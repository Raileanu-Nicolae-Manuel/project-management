package api

import (
	"backend/internal/db"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	users, err := queries.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter out sensitive information
	var response []userResponse
	for _, user := range users {
		response = append(response, userResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := queries.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(30 * 24 * time.Hour).Unix(),
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	// Set the token in the response
	w.Header().Set("Authorization", "Bearer "+tokenString)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func Register(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Create user
	params := db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	result, err := queries.CreateUser(r.Context(), params)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "User registered successfully",
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := queries.DeleteUser(r.Context(), id); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

func UpdateUser(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := db.UpdateUserParams{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
	}

	if err := queries.UpdateUser(r.Context(), params); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func RegisterRoutes(handlerRouter *HandlerRouter, path string) {
	handlerRouter.router.Route(path, func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			GetUsers(w, r, handlerRouter.queries)
		})
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			Register(w, r, handlerRouter.queries)
		})
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			Login(w, r, handlerRouter.queries)
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			UpdateUser(w, r, handlerRouter.queries)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			DeleteUser(w, r, handlerRouter.queries)
		})
	})
}
