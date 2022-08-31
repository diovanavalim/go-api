package controller

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.Validate("create"); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	} else {
		if err := user.Format("create"); err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			return
		}
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
	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repository.CreateUserRepository(db)

	result, err := userRepository.GetUsers()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func GetUsersByNameOrNickname(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repository.CreateUserRepository(db)

	result, err := userRepository.GetUsersByNameOrNickname(nameOrNickname)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repository.CreateUserRepository(db)

	result, err := userRepository.GetUserByID(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	requestUserID, err := auth.ExtractUserID(r)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if requestUserID != userId {
		response.Error(w, http.StatusForbidden, errors.New("Request user's ID does not match user ID query param"))
		return
	}

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

	if err = user.Validate("update"); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	} else {
		if err := user.Format("update"); err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repository.CreateUserRepository(db)

	if err = userRepository.UpdateUser(userId, user); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	requestUserID, err := auth.ExtractUserID(r)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if requestUserID != userId {
		response.Error(w, http.StatusForbidden, errors.New("Request user's ID does not match user ID query param"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repository.CreateUserRepository(db)

	if err = userRepository.DeleteUser(userId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
