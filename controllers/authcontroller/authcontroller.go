package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/jeypac/go-jwt-mux/helper"
	"github.com/jeypac/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// GET INPUT JSON

	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	// HASH PASSWORD
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if userInput.Password == "" {
		response := map[string]string{"message": "Password cannot be empty"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	userInput.Password = string(hashPassword)

	// INSERT DATABASE
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
