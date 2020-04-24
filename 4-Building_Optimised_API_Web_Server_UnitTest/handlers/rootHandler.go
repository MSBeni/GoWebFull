package handlers

import (
	"fmt"
	"net/http"
)
// RootHandler handles the root
func RootHandler(w http.ResponseWriter, r *http.Request){
	// We add this if statement to mention that the root path is explicitly the "/"
	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The Access is not Defined \n"))
		return
	}
	// if the path is correct then the code will be written to the server
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RUNNING API VERSION 1 \n"))
	fmt.Fprint(w, "The easiest API \n")

}