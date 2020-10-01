package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	//"sort"
	//"strconv"

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
		Authentication: driver.BasicAuthentication("root", "...."),
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

	ctx := context.Background()

	getUser := fmt.Sprintf(`FOR u IN Users
									  FOR f IN Friendship
										FILTER u._id == f._to
										RETURN u.UserName`)

	//bindVars := map[string]interface{}{
	//	"_id": _id,
	//}
	cursor, err := db.Query(ctx, getUser, nil)
	if err != nil {
		panic(err)
	}
	//defer cursor.Close()

	if !cursor.HasMore() {
		panic(err)
	}
	//OutgoingList := []UsersNew{}

	users := []string{}

	for {
		//user_ := UsersNew{}
		var usernames string
		_, err := cursor.ReadDocument(ctx, &usernames)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
		//fmt.Println(user_)
		//OutgoingList = append(OutgoingList, usernames)
		users = append(users, usernames)
	}
	uniqueUsernames := []string{}
	//fmt.Println(users)
	for i, el := range users{
		//fmt.Println(i)
		if i == 0 {
			uniqueUsernames = append(uniqueUsernames, el)
		}
		n := 0
		for _, UniqVal := range uniqueUsernames{
			if el == UniqVal{
				n += 1
			}
		}
		if n == 0{
			uniqueUsernames = append(uniqueUsernames, el)
		}
	}
	//fmt.Println(uniqueUsernames)
	//fmt.Println(len(uniqueUsernames))
	UsersRep := map[string]int{}
	for _, Newel := range uniqueUsernames{
		Repet := 0
		for _, RepEl := range users{
			if Newel == RepEl {
				Repet += 1
			}
		}
		UsersRep[Newel] = Repet
	}


	chosenUsers := map[string]int{}
	for key, val := range UsersRep{
		if val >= 20{
			chosenUsers[key] = val
		}
	}
	OutgoingList := []UsersNew{}

	for usn, _:= range chosenUsers{
		getUserbyUsername := fmt.Sprintf(`FOR user IN Users FILTER user.%s == @UserName RETURN user`, UserName)
		bindVars := map[string]interface{}{
			"UserName": usn,
		}
		cursor, err := db.Query(ctx, getUserbyUsername, bindVars)
		if err != nil {
			panic(err)
		}
		//defer cursor.Close()

		if !cursor.HasMore() {
			panic(err)
		}
		user_ := UsersNew{}

		_, err = cursor.ReadDocument(ctx, &user_)

		OutgoingList = append(OutgoingList, user_)
	}
	fmt.Println(OutgoingList)
	//for _, k := range keys {
	//	fmt.Println(k, UsersRep[k])
	//}
	//
	//fmt.Println(UsersRep[k])
	//var user UsersNew
	//meta, err := cursor.ReadDocument(ctx, &user)
	//fmt.Println("Who is Followed is >>>>> ", meta.ID, user)

}
