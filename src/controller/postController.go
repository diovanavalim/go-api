package controller

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post model.Post

	if err = json.Unmarshal(requestBody, &post); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	post.Format()

	userId, err := auth.ExtractUserID(r)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	postRepository := repository.CreateRepository(db)

	result, err := postRepository.CreatePost(post, userId)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, struct {
		PostID uint64 `json:"postId"`
	}{
		PostID: result,
	})

}

func GetPosts(w http.ResponseWriter, r *http.Request) {}

func GetPost(w http.ResponseWriter, r *http.Request) {}

func UpdatePost(w http.ResponseWriter, r *http.Request) {}

func DeletePost(w http.ResponseWriter, r *http.Request) {}
