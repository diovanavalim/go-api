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

	"github.com/gorilla/mux"
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

	postRepository := repository.CreatePostRepository(db)

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

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	postRepository := repository.CreatePostRepository(db)

	result, err := postRepository.GetPosts(userID)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	postRepository := repository.CreatePostRepository(db)

	result, err := postRepository.GetPost(postId)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userID, err := auth.ExtractUserID(r)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	postRepository := repository.CreatePostRepository(db)

	post, err := postRepository.GetPost(postID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if post.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("User shall not delete a post from other user"))
		return
	}

	if err = postRepository.DeletePost(postID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
