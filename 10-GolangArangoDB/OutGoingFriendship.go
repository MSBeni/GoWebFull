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
		Authentication: driver.BasicAuthentication("root", "root1!!!"),
	})
	if err != nil {
		panic(err)

	}
	// Open "examples_books" database
	db, err := c.Database(nil, "_system")
	if err != nil {
		panic(err)
	}
	_id := "Users/f7755fa8"
	//Users := "Users"
	//_from := "_from"
	//_to := "_to"
	//NEWAMOUNT := "387"
	//
	//_FromUserName := "AmberMonarrez"
	//_ToUserName := "JunitaBrideau"
	ctx := context.Background()

	getUser := fmt.Sprintf(`FOR v, e, p in 1..1 OUTBOUND @_id GRAPH 'Followings' RETURN v`)

	bindVars := map[string]interface{}{
		"_id": _id,
	}
	cursor, err := db.Query(ctx, getUser, bindVars)
	if err != nil {
		panic(err)
	}
	//defer cursor.Close()

	if !cursor.HasMore() {
		panic(err)
	}
	OutgoingList := []UsersNew{}

	for {
		user_ := UsersNew{}

		_, err := cursor.ReadDocument(ctx, &user_)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
		//fmt.Println(user_)
		OutgoingList = append(OutgoingList, user_)
	}
	fmt.Println(OutgoingList)
	//var user UsersNew
	//meta, err := cursor.ReadDocument(ctx, &user)
	//fmt.Println("Who is Followed is >>>>> ", meta.ID, user)

}
