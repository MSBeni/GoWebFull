package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	//"log"
	//"time"

	//"strconv"
	//"time"
)


type UsersNew struct {
	Key       string `json:"_key,omitempty"`
	UserName  string `json:"UserName"`
	LowerName string `json:"LowerName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	// Email is the primary email address (to be used for communication)
	Email            string `json:"Email"`
	KeepEmailPrivate string   `json:"KeepEmailPrivate"`
	PasswordHash     string `json:"PasswordHash"`
	PasswordCost     string    `json:"PasswordCost"`
	PasswordHashAlg  string `json:"PasswordHashAlg"`

	// MustChangePassword is an attribute that determines if a user
	// is to change his/her password after registration.
	MustChangePassword string   `json:"MustChangePassword"`
	Location           string `json:"Location"`
	Language           string `json:"Language"`
	Description        string `json:"Description"`

	CreatedUnix   string `json:"CreatedUnix"`
	UpdatedUnix   string `json:"UpdatedUnix"`
	LastLoginUnix string `json:"LastLoginUnix"`

	// Permissions
	IsActive     string `json:"IsActive"` // Activate primary email
	IsAdmin      string `json:"IsAdmin"`
	IsVerified   string `json:"IsVerified"`
	IsRestricted string `json:"IsRestricted"`

	// Avatar
	Avatar      string `json:"Avatar"`
	AvatarEmail string `json:"AvatarEmail"`

	// Counters
	NumFollowers string `json:"NumFollowers"`
	NumFollowing string `json:"NumFollowing"`
	NumStars     string `json:"NumStars"`
	NumRepos     string `json:"NumRepos"`
}

type Assets struct {
	AccountID string
	Amount string
}

type Friendship struct {
	User UsersNew
	Created string
	Asset Assets
	//// Reference to another document. Format: ':collection/:key'
	//From string `json:"_from,omitempty"`
	//// Reference to another document. Format: ':collection/:key'
	//To string `json:"_to,omitempty"`
}
func main() {

	conn, err := http.NewConnection(http.ConnectionConfig{
	Endpoints: []string{"http://localhost:8529"},
	TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
	fmt.Println(err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
	Connection:     conn,
	Authentication: driver.BasicAuthentication("root", "..."),
	})
	if err != nil {
	panic(err)

	}
	// Open "examples_books" database
	db, err := c.Database(nil, "_system")
	if err != nil {
	panic(err)
	}
	UserName := "UserName"
	Users := "Users"
	_from := "_from"
	_to := "_to"
	_FromUserName := "AmberMonarrez"
	_ToUserName := "JunitaBrideau"
	ctx := context.Background()

	getUser := fmt.Sprintf(`FOR user IN %s FILTER user.%s == @UserName RETURN user`, Users, UserName)

	bindVars := map[string]interface{}{
		"UserName": _ToUserName,
	}
	cursor, err := db.Query(ctx, getUser, bindVars)
	if err != nil {
		panic(err)
	}
	//defer cursor.Close()

	if !cursor.HasMore() {
		panic(err)
	}

	var user UsersNew
	meta, err := cursor.ReadDocument(ctx, &user)
	fmt.Println("Who is Followed is >>>>> ", meta.ID, user)

	bindVars2 := map[string]interface{}{
		"UserName": _FromUserName,
	}
	cursor2, err := db.Query(ctx, getUser, bindVars2)
	if err != nil {
		panic(err)
	}
	//defer cursor.Close()

	if !cursor2.HasMore() {
		panic(err)
	}

	var user2 UsersNew
	meta2, err := cursor2.ReadDocument(ctx, &user2)
	fmt.Println("Follower is >>>>> ", meta2.ID, user2)


	DeleteFriendship := fmt.Sprintf(`FOR user IN Friendship FILTER user.%s == @Follower && user.%s == @Following REMOVE user IN Friendship`, _from, _to)
	bindVars3 := map[string]interface{}{
		"Follower": meta2.ID,
		"Following": meta.ID,
	}
	cursor, err = db.Query(ctx, DeleteFriendship, bindVars3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Friendship deleting is successfully Done")
}
