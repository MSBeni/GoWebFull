package users

import (
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M){
	m.Run()
	os.Remove(dbPath)
}

func TestCRUD(t *testing.T){
	t.Log("Create")
	u := &User{
		ID:   bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}
	err := u.Save()
	if err!=nil{
		log.Fatalf("Error saving the record %s", err)
	}
	t.Log("Read")
	u2, err := One(u.ID)
	if err!=nil{
		log.Fatalf("Error retreiving the record %s", err)
	}
	if !reflect.DeepEqual(u2, u){
		t.Error("Record do not match")
	}
	t.Log("Update")
	u.Role = "developer"
	err = u.Save()
	if err!=nil{
		log.Fatalf("Error saving the record %s", err)
	}
	u3, err := One(u.ID)
	if err!=nil{
		log.Fatalf("Error retreiving the record %s", err)
	}
	if !reflect.DeepEqual(u3, u){
		t.Error("Record do not match")
	}
	t.Log("Delete")
	err = Delete(u.ID)
	if err!=nil{
		log.Fatalf("Error removing the record %s", err)
	}
	_, err = One(u.ID)
	if err == nil{
		t.Fatal("Error should not exist anymore")
	}
	if err != storm.ErrNotFound{
		t.Fatalf("Error retreiving non-existing record %s", err)
	}
	t.Log("Read All")
	u2.ID = bson.NewObjectId()
	u3.ID = bson.NewObjectId()
	err = u2.Save()
	if err!=nil{
		log.Fatalf("Error saving the record %s", err)
	}
	err = u3.Save()
	if err!=nil{
		log.Fatalf("Error saving the record %s", err)
	}
	users, err := All()
	if err!= nil{
		log.Fatalf("Error reading all records %s", err)
	}
	if len(users) != 4{
		t.Errorf("Different number of records retrieved. Expected: 3 / Actual: %d ", len(users))

	}
}