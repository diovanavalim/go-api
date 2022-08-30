package controller

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user model.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.CreateUserRepository(db)

	result, err := userRepository.CreateUser(user)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User successfully inserted, ID: %d", result)))
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
