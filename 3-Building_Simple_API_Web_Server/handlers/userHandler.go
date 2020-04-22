package handlers

import (
	"net/http"
	"GoWebFull/3-Building_Simple_API_Web_Server/users"
)
func usersGetAll(w http.ResponseWriter, r *http.Request){
	users, err := users.All()
	if err != nil{
		PostError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})

}
