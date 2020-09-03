package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
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
	// Reference to another document. Format: ':collection/:key'
	From string `json:"_from,omitempty"`
	// Reference to another document. Format: ':collection/:key'
	To string `json:"_to,omitempty"`
}

func createEdgeDefinitions(edgeCollections []string, fromVertices []string, toVertices []string) []driver.EdgeDefinition {
	edgeDefinitions := []driver.EdgeDefinition{}
	for _, edge := range edgeCollections {
		var definition driver.EdgeDefinition
		definition.Collection = edge
		definition.From = fromVertices
		definition.To = toVertices
		edgeDefinitions = append(edgeDefinitions, definition)
	}

	return edgeDefinitions
}
//
//func createUserGraph(ctx context.Context, userID string, eColName string, dColName string, db driver.Database) (driver.Graph, error) {
//	gopt := &driver.CreateGraphOptions{
//		EdgeDefinitions: []driver.EdgeDefinition{driver.EdgeDefinition{
//			Collection: eColName,
//			From:       []string{dColName},
//			To:         []string{dColName},
//		}},
//	}
//	return db.CreateGraph(ctx, userID, gopt)
//}



func main(){

	UserToCheck := "08dfd236"
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		fmt.Println(err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
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
	ctx := context.Background()

	// check if the graph exist
	var graph driver.Graph
	if ok, _ := db.GraphExists(nil, "Followings"); ok {
		graph, _ = db.Graph(nil, "Followings")
	} else {
		graph, _ = db.CreateGraph(nil, "Followings", nil)
	}
	fmt.Println("The available graph is: ", graph.Name())

	query2 := "FOR v, e, p in 1..1 INBOUND 'Users/e83d1009' GRAPH 'Followings' RETURN v"

	cursor, err := db.Query(ctx, query2, nil)
	if err != nil {
		panic(err)
	}
	FollowersNum := 0
	for {
		relation_ := Friendship{}

		meta, err := cursor.ReadDocument(ctx, &relation_)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}

		//fmt.Printf("The id of the relation is >>>>>> '%s' '\n'" , relation_.User)
		fmt.Println("The key is >>>>>>  ", meta.Key)
		if meta.Key == UserToCheck {
			fmt.Printf("The friendship already exist between %s and %s", "e83d1009", UserToCheck)
			FollowersNum += 1
		}
	}
	if FollowersNum == 0{
		fmt.Println("Ready to submit new friendship")
	}


	query3 := "FOR v, e, p in 1..1 INBOUND 'Users/e83d1009' GRAPH 'Followings' FILTER v._id == 'Users/08dfd236' RETURN e"
	cursor4, err := db.Query(ctx, query3, nil)
	for {
		relation_ := Friendship{}
		meta, err := cursor4.ReadDocument(ctx, &relation_)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("New Added Follower is >>>>>> %s \n", relation_.User)
		fmt.Printf("New Added Follower is >>>>>> %s \n", meta.ID)

	}

	
}