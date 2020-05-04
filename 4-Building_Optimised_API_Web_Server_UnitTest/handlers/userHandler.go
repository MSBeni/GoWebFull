package handlers

import (
	"GoWebFull/4-Building_Optimised_API_Web_Server_UnitTest/cache"
	"GoWebFull/4-Building_Optimised_API_Web_Server_UnitTest/users"
	"encoding/json"
	"errors"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

func bodyToUser(r *http.Request, u *users.User) error{
	if r == nil{
		return errors.New("a request is required")
	}
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
	if cache.Serve(w, r){
		return
	}
	users, err := users.All()
	if err != nil{
		PostError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead{
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"users": users})

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
	cache.Drop("/users")
	w.Header().Set("Location", "/users/" + u.ID.Hex())
	// write the code to the header
	w.WriteHeader(http.StatusCreated)
}

func usersGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId){
	if cache.Serve(w, r){
		return
	}
	u, err := users.One(id)
	 if err != nil{
		 if err == storm.ErrNotFound{
		 	// StatusNotFound = 404 // RFC 7231, 6.5.4
		 	PostError(w, http.StatusNotFound)
		 	return
		 }
		 // StatusInternalServerError = 500 // RFC 7231, 6.6.1
		 PostError(w, http.StatusInternalServerError)
		 return
	 }
	if r.Method == http.MethodHead{
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	 // StatusOK = 200 // RFC 7231, 6.3.1
	cw := cache.NewWriter(w, r)
	 postBodyResponse(cw, http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId){
	u := new(users.User)
	err := bodyToUser(r, u)
	if err != nil{
		PostError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil{
		if err == users.ErrRecordInvalid{
			PostError(w, http.StatusBadRequest)
		}else{
			PostError(w, http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"user": u})
}

func usersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId){
	u, err := users.One(id)
	if err != nil{
		if err == storm.ErrNotFound{
			// StatusNotFound = 404 // RFC 7231, 6.5.4
			PostError(w, http.StatusNotFound)
			return
		}
		// StatusInternalServerError = 500 // RFC 7231, 6.6.1
		PostError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToUser(r, u)
	if err != nil{
		PostError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	err = u.Save()
	if err != nil{
		if err == users.ErrRecordInvalid{
			PostError(w, http.StatusBadRequest)
		}else{
			PostError(w, http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	cw := cache.NewWriter(w, r)
	postBodyResponse(cw, http.StatusOK, jsonResponse{"user": u})
}

func usersDeleteOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId){
	err := users.Delete(id)
	if err != nil{
		if err == storm.ErrNotFound{
			// StatusNotFound = 404 // RFC 7231, 6.5.4
			PostError(w, http.StatusNotFound)
			return
		}
		// StatusInternalServerError = 500 // RFC 7231, 6.6.1
		PostError(w, http.StatusInternalServerError)
		return
	}
	cache.Drop("/users")
	cache.Drop(cache.MakeResource(r))
	w.WriteHeader(http.StatusOK )
}

