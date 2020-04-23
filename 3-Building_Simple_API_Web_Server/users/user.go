
package users

import (
	"errors"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// User type in order to define a user
type User struct {
	ID   bson.ObjectId  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

const (
	dbPath = "users.db"
)

// errors
var(
	ErrRecordInvalid = errors.New("The record is invalid")
)

// All retrieves all users from the database
func All()([]User, error){
	db, err := storm.Open(dbPath)
	if err != nil{
		return nil, err
	}
	// it's necessary to close the opened database to prevent data leakage and other necessary issues
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil{
		return nil, err
	}
	return users, nil
}

// One returns a single user from the database with the given id
func One(id bson.ObjectId) (*User, error){
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)

	err = db.One("ID", id, u)
	if err != nil{
		return nil, err
	}
	return u, nil
}

// Delete is a function to delete the record with the mentioned id from the database
func Delete(id bson.ObjectId) error{
	db, err := storm.Open(dbPath)
	if err != nil{
		return err
	}
	defer db.Close()
	s := new(User)
	err = db.One("ID", id, s)
	if err != nil{
		return err
	}
	return db.DeleteStruct(s)
}

// Save the updates or create the given record in the database
func (u *User) Save() error{
	if err := u.Validate(); err != nil{
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil{
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// Validate the user with the given username
func (u *User) Validate() error{
	if u.Name == ""{
		return ErrRecordInvalid
	}
	return nil
}