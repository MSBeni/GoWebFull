package handlers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

// handle the user route
func UsersRouter(w http.ResponseWriter, r *http.Request){
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users"{
		switch r.Method {
		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			usersPostOne(w, r)
			return
		default:
			PostError(w, http.StatusMethodNotAllowed)

		}
	}

	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path){
		PostError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		usersGetOne(w, r, id)
		return
	case http.MethodPut:
		return
	case http.MethodDelete:
		return
	case http.MethodPatch:
		return
	default:
		PostError(w, http.StatusMethodNotAllowed)

	}


}