package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"time"

	//"strconv"
	//"time"
)

type Users struct {
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
	User Users
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
		Authentication: driver.BasicAuthentication("...", "..."),
	})
	if err != nil {
		panic(err)

	}
	// Open "examples_books" database
	db, err := c.Database(nil, "_system")
	if err != nil {
		panic(err)
	}

	_from := "_from"
	_to := "_to"

	ctx := context.Background()
	query := fmt.Sprintf(`FOR user IN Friendship FILTER user.%s == @Follower && user.%s == @Following RETURN user._id`, _from, _to)

	//query := fmt.Sprintf(`FOR user IN Friendship FILTER user._from == "Users/08dfd236" && user._to == "Users/e83d1009" RETURN user`)
	bindVars := map[string]interface{}{
		"Follower": "Users/f7755fa8",
		"Following": "Users/08dfd236",
	}
	cursor, err := db.Query(ctx, query, bindVars)

	//cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal(err)
	}


	var id string


	_, err = cursor.ReadDocument(ctx, &id)
	if err != nil{
		query2 := fmt.Sprintf(`FOR user IN Users FILTER user._id == @id RETURN user`)
		bindVars2 := map[string]interface{}{
			"id": "Users/f7755fa8",
		}
		cursor2, err := db.Query(ctx, query2, bindVars2)
		var user Users
		_, err = cursor2.ReadDocument(ctx, &user)
		fmt.Println(">>>>>>>>>>>>>", user)

		querynew := fmt.Sprintf(`INSERT { _from: @Follower, _to: @Following, User:  @UserFollower, Created:@CreateTime, Asset:@Assets} INTO Friendship`)
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(">>>>>>>>>>>>>", timeNow)
		bindVars = map[string]interface{}{
			"Follower": "Users/f7755fa8",
			"Following": "Users/08dfd236",
			"UserFollower": user,
			"CreateTime":timeNow,
			"Assets":Assets{"z2e09", "154"},
		}
		cursor, err = db.Query(ctx, querynew, bindVars)
		if err != nil{
			panic(err)
		}

	}else{
		fmt.Println("ID of the Friendship is >>>>>", id)
	}


	//fmt.Println("ID of the Friendship is >>>>>", id)

}




