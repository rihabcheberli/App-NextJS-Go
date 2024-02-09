package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rihabcheberli/App-NextJS-Go/backend/database"
	"github.com/rihabcheberli/App-NextJS-Go/backend/models"
	"github.com/rihabcheberli/App-NextJS-Go/backend/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ValidateUser(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user input: "+err.Error())
		return
	}

	_, err := database.GetUserByEmail(user.Email)
	if err == nil {
		utils.RespondWithError(w, http.StatusConflict, "User already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user.Password = hashedPassword

	if err := database.CreateUser(&user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user account: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	storedUser, err := database.GetUserByEmail(user.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if !utils.ComparePasswords(user.Password, storedUser.Password) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Login successful"})
}
