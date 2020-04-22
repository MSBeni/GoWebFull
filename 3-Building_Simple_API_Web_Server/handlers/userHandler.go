package handlers

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"GoWebFull/3-Building_Simple_API_Web_Server/users"
)

func bodyToUser(r *http.Request, u *users.User) error{
	if r.Body == nil{
		return errors.New("request body is empty")
	}
	if u == nil{
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}
	return json.Unmarshal(bd, u)
}


func usersGetAll(w http.ResponseWriter, r *http.Request){
	users, err := users.All()
	if err != nil{
		PostError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})

}

func usersPostOne(w http.ResponseWriter, r *http.Request){
	u := new(users.User)
	err := bodyToUser(r, u)
	if err != nil{
		PostError(w, http.StatusBadRequest)
		return
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil{
		if err == users.ErrRecordInvalid{
			PostError(w, http.StatusBadRequest)
		}else{
			PostError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/users/" + u.ID.Hex())
	// write the code to the header
	w.WriteHeader(http.StatusCreated)
}
