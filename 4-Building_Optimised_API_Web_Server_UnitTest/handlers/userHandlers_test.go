package handlers

import (
	"GoWebFull/4-Building_Optimised_API_Web_Server_UnitTest/users"
	"bytes"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestBodyByUser(t *testing.T){
	valid := &users.User{
		ID: bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}
	valid2 := &users.User{
		ID: valid.ID,
		Name: "John",
		Role: "Developer",
	}
	js, err := json.Marshal(valid)
	if err != nil{
		t.Errorf("Error Marshaling a valid user: %s", err)
		t.FailNow()
	}
	ts := []struct{
		txt string            // to describe the test case
		r *http.Request       // first argument to the function
		u *users.User          // second argument to the function
		err bool              // to mark is an error is expected
		exp *users.User        // the data that we expect to contain after the function is run
	}{
		{
			txt: "nil request",
		 	err: true,
		},
		{
		txt: "empty request body",
		r: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
		},
		err: true,
		},
		{
		txt: "malformed data",
		r: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":12}`)),
		},
		u: &users.User{},
		err: true,
		},
		{
		txt: "valid request",
		r: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBuffer(js)),
		},
		u: &users.User{},
		exp: valid,
		},
		{
		txt: "valid partial request",
		r: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"role":"Developer", "Age":22}`)),
		},
		u: valid,
		exp: valid2,
		},

	}
	for _, tc := range ts{
		t.Log(tc.txt)      // by logging the test text we make sure only the failing tests displays this information
		err := bodyToUser(tc.r, tc.u)
		if tc.err{
			if err == nil{
				t.Error("Expected error, got none.")
			}
			continue
		}
		if err != nil{
			t.Errorf("Unexpected error: %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.u, tc.exp){
			t.Error("Unmarshalled stat is different:")
			t.Error(tc.u)
			t.Error(tc.exp)
		}
	}
}