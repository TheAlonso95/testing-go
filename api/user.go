package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/goschool/crud/db"
	"github.com/goschool/crud/types"
	"github.com/goschool/crud/utils"
)

type UserHandler struct {
	userStore db.UserStore
}

type LoginResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(us db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: us,
	}
}

func (u *UserHandler) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	var createUser types.CreateUser
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	user, err := NewUser(createUser)
	if err != nil {
		http.Error(w, "Invalid New user", http.StatusBadRequest)
		return
	}

	newUser, err := u.userStore.CreateUser(ctx, user)

	log.Print(newUser)
	if err != nil {
		http.Error(w, "Invalid New user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "Unable to register new user", http.StatusInternalServerError)
	}
}

func (u *UserHandler) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {
	var loginParams LoginParams
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&loginParams); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}

	user, _ := u.userStore.GetUser(ctx, loginParams.Email)

	isPasswordCorrect := utils.ValidatePassword(loginParams.Password, user.PasswordHash)
	if !isPasswordCorrect {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	token, err := utils.CreateToken(*user)
	if err != nil {
		// Not sure if that's the correct approach here!
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	loginResponse := &LoginResponse{
		User:  user,
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(*loginResponse); err != nil {
		http.Error(w, "Unable to register new user", http.StatusInternalServerError)
	}
}

func NewUser(createUser types.CreateUser) (*types.User, error) {
	hashPassword, err := utils.HashPassword(createUser.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to create password hash: %w", err)
	}

	return &types.User{
		Name:         createUser.Name,
		Email:        createUser.Email,
		PasswordHash: string(hashPassword),
	}, nil
}
