package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
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

func main(){

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		fmt.Println(err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication("...", "..."),
	})
	if err != nil {
		panic(err)

	}
	// Open "examples_books" database
	db, err := c.Database(nil, "Comvest")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	query := "FOR d IN Friendship LIMIT 10 RETURN d"
	//query :="FOR doc IN Friendship FILTER doc.foo == @BrittniGillaspie RETURN doc"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.Background()
	getUsertbyID2 := fmt.Sprintf(`FOR user IN Users FILTER user.FirstName == @name RETURN user`)
	//query3 := `FOR user IN Users FILTER user._key == @key RETURN user`
	bindVars := map[string]interface{}{
		"name": "James",
	}
	//bindVars2 := map[string]interface{}{
	//	"key": "3e5e6a05",
	//}
	cursor, err = db.Query(ctx, getUsertbyID2, bindVars)
	//cursor, err = db.Query(ctx, query3, bindVars2)
	if err != nil {
		panic(err)
	}
	//defer cursor.Close()
	for {
		user_ := Users{}

		_, err := cursor.ReadDocument(ctx, &user_)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
		fmt.Println(user_)
	}

	var graph driver.Graph
	if ok, _ := db.GraphExists(nil, "Followings"); ok {
		graph, _ = db.Graph(nil, "Followings")
	} else {
		graph, _ = db.CreateGraph(nil, "Followings", nil)
	}


	fmt.Println(graph.Name())

	fmt.Println(graph.EdgeCollections(ctx))


	query2 := "FOR v, e, p in 1..1 INBOUND 'Users/e83d1009' GRAPH 'Followings' RETURN e"

	cursor, err = db.Query(ctx, query2, nil)
	if err != nil {
		panic(err)
	}

	for {
		relation_ := Friendship{}

		meta, err := cursor.ReadDocument(ctx, &relation_)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		query4 := `FOR user IN Users FILTER user._key == @key RETURN user`
		bindVars3 := map[string]interface{}{
			"key": meta.Key,
		}
		fmt.Printf("###### the Key of the user is  >>>>>>>>> '%s' '\n'", meta.Key)
		//cursor, err = db.Query(ctx, getUsertbyID2, bindVars)
		cursor2, err := db.Query(ctx, query4, bindVars3)
		if err != nil {
			panic(err)
		}
		//defer cursor.Close()
		for {
			user_ := Users{}

			_, err := cursor2.ReadDocument(ctx, &user_)
			if driver.IsNoMoreDocuments(err) {
				break
			} else if err != nil {
				panic(err)
			}
			//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
			fmt.Println(user_)
		}
	}






}
