package handlers

import (
	"encoding/json"       // Package json implements encoding and decoding of JSON as defined in RFC 7159.
	"net/http"
)


type jsonResponse map[string]interface{}

func PostError(w http.ResponseWriter, code int){
	http.Error(w, http.StatusText(code), code)
}


func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse){
	if content !=nil{
		//func Marshal(v interface{}) ([]byte, error)
		//Marshal returns the JSON encoding of v.
		//Marshal traverses the value v recursively. If an encountered value implements the Marshaler interface and is
		//not a nil pointer, Marshal calls its MarshalJSON method to produce JSON
		js, err := json.Marshal(content)
		if err != nil{
			PostError(w, http.StatusInternalServerError)
		}
		// Header returns the header map that will be sent by
		// WriteHeader. The Header map also is the mechanism with which
		// Handlers can set HTTP trailers.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return
	}
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
}