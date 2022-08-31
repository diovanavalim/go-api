package controller

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"api/src/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	userRepository := repository.CreateUserRepository(db)

	result, err := userRepository.GetUserByEmail(user.Email)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := util.CompareHash(result.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	response.JSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "Successfully logged in",
		Token:   "TODO",
	})
}
