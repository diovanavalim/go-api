package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating New User!"))
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
