package users

import (
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestMain(m *testing.M){
	m.Run()
	os.Remove(dbPath)
}

func cleanDb(b *testing.B) {
	os.Remove(dbPath)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "John",
			Role: "Tester",
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving the record %s", err)
		}
	}
	b.ResetTimer()
}

func BenchmarkCreate(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "John" + strconv.Itoa(i),
			Role: "Tester",
		}
		b.StartTimer()
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving the record %s", err)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "John" + strconv.Itoa(i),
			Role: "Tester",
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving the record %s", err)
		}
		b.StartTimer()
		_, err = One(u.ID)
		if err!=nil{
			b.Fatalf("Error retreiving the record %s", err)
		}
	}
}

func BenchmarkUpdate(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "John" + strconv.Itoa(i),
			Role: "Tester",
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving the record %s", err)
		}
		b.StartTimer()
		u.Role = "developer"
		err = u.Save()
		if err!=nil{
			b.Fatalf("Error saving the record %s", err)
		}

	}
}


func BenchmarkDelete(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "John" + strconv.Itoa(i),
			Role: "Tester",
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving the record %s", err)
		}
		b.StartTimer()
		err = Delete(u.ID)
		if err!=nil{
			b.Fatalf("Error removing the record %s", err)
		}
	}
}

func BenchmarkCRUD(b *testing.B){
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i<b.N;i++{
	u := &User{
		ID:   bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}
	err := u.Save()
	if err!=nil{
		b.Fatalf("Error saving the record %s", err)
	}
	_, err = One(u.ID)
	if err!=nil{
		b.Fatalf("Error retreiving the record %s", err)
	}
	u.Role = "developer"
	err = u.Save()
	if err!=nil{
		b.Fatalf("Error saving the record %s", err)
	}
	err = Delete(u.ID)
	if err!=nil{
		b.Fatalf("Error removing the record %s", err)
	}
	}
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
	if len(users) != 2{
		t.Errorf("Different number of records retrieved. Expected: 3 / Actual: %d ", len(users))

	}
}