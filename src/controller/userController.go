package controller

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	} else {
		user.Format()
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repository.CreateUserRepository(db)

	result, err := userRepository.CreateUser(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
		UserID  uint64 `json:"user_id"`
	}{
		Message: "User successfully inserted",
		UserID:  result,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Users!"))
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Specific User!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User!"))
}
