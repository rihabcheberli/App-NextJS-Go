package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rihabcheberli/App-NextJS-Go/backend/database"
	"github.com/rihabcheberli/App-NextJS-Go/backend/models"
	"github.com/rihabcheberli/App-NextJS-Go/backend/utils"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	user, err := database.GetUserByID(userID)
	if err != nil {
		if err == database.ErrUserNotFound {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve user data: "+err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	if err := ValidateUser(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user input: "+err.Error())
		return
	}

	err = database.CreateUser(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user account: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	if err := ValidateUser(updatedUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user input: "+err.Error())
		return
	}

	err = database.UpdateUserByID(userID, updatedUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user account: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	err := database.DeleteUserByID(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user account: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch users from the database")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func ValidateUser(user models.User) error {
	if user.Email == "" {
		return errors.New("Email is required")
	}

	return nil
}
