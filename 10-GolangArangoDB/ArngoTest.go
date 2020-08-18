package main

import (
	"context"
	"strconv"

	//"bitbucket.org/comvest-services/pkg/api/common"
	"crypto/tls"
	"fmt"
	//"reflect"
	//"crypto/tls"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"GoWebFull/10-GolangArangoDB/pkg/query"
	//client "github.com/influxdata/influxdb1-client/v2"
	"GoWebFull/10-GolangArangoDB/pkg/errors"
)

type MyDocument struct{
	Symbol string `json:"Symbol"`
	Avatar string `json:"Avatar"`
}
const (
	databaseName          = "Comvest"
	productCollectionName = "ProductList"
)

type Service struct {
	dbClient   driver.Client
}

// Products represents the object of Cryptocurrencies
type Product struct {
	Key       string `json:"_key,omitempty"`
	BaseCurrency  string `json:"BaseCurrency"`
	QuoteCurrency string `json:"QuoteCurrency"`
	BaseMinSize string `json:"BaseMinSize"`
	BaseMaxSize  string `json:"BaseMaxSize"`
	QuoteIncrement  string `json:"QuoteIncrement"`
}



func (s *Service) Search(offset int, count int) (interface{}, int64, error) {
	result, resultCount, err := s.searchProduct(offset, count)
	if err != nil {
		return nil, 0, err
	}
	return result, resultCount, nil
}

func (s *Service) searchProduct(offset int, count int) (interface{}, int64, error) {

	queryBuilder := query.NewForQuery(productCollectionName, "doc")
	qs := queryBuilder.
		LIMIT(offset, count).
		Return().String()

	pCtx := context.Background()
	ctx := driver.WithQueryCount(pCtx)

	// Open database
	db, err := s.dbClient.Database(pCtx, databaseName)
	if err != nil {
		return nil, 0, errors.New(err, fmt.Sprintf("failed to open [%s] database", databaseName))
	}

	cursor, err := db.Query(ctx, qs, nil)
	if err != nil {
		return nil, 0, errors.New(err, fmt.Sprintf("failed to query [%s] on %s database", qs, databaseName))
	}

	defer cursor.Close()

	result := []MyDocument{}
	if cursor.Count() == 0 || !cursor.HasMore() {
		return result, cursor.Count(), nil
	}

	for {
		product := MyDocument{}
		_, err = cursor.ReadDocument(pCtx, &product)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, errors.New(err, fmt.Sprintf("failed to read response from cursor for query [%s]", qs))
		}

		result = append(result, product)
	}

	return result, cursor.Count(), nil
}


func main(){

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		// Handle error
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication("...", "..."),
	})
	if err != nil {
		// Handle error

	}
	// Open "examples_books" database
	db, err := c.Database(nil, "Comvest")
	if err != nil {
		// Handle error
	}

	//// Open "books" collection
	//col, err := db.Collection(nil, "flights")
	//if err != nil {
	//	// Handle error
	//}
	//
	//if err != nil {
	//	// Handle error
	//}
	//fmt.Printf("Created document in collection '%s' in database '%s'\n", col.Name(), db.Name())

	//var doc MyDocument
	//ctx := context.Background()
	////myDocumentKey := "3808877"
	//myDocumentKey := "319"
	//meta, err := col.ReadDocument(ctx, myDocumentKey, &doc)
	//if err != nil {
	//	// handle error
	//}
	//fmt.Println(meta)
	//
	//ctx = driver.WithQueryCount(context.Background())
	//
	//query := "FOR d IN ProductsList RETURN d"
	//cursor, err := db.Query(ctx, query, nil)
	//if err != nil {
	//	// handle error
	//}
	//defer cursor.Close()
	////fmt.Printf("Query yields %d documents\n", cursor.Count())
	////fmt.Println(cursor.Statistics())
	//
	//ctx = context.Background()
	//query = "FOR d IN ProductsList FILTER d.Name == @name RETURN d"
	//bindVars := map[string]interface{}{
	//	"name": "Some name",
	//}
	//cursor, err = db.Query(ctx, query, bindVars)
	//if err != nil {
	//	// handle error
	//}
	//defer cursor.Close()


	ctx := context.Background()
	query := "FOR d IN ProductsList2 LIMIT 10 RETURN d"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer cursor.Close()
	var result []string
	for {
		var doc MyDocument

		meta, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		fmt.Printf("Got doc with key '%s' from query\n", meta.Key)
		//fmt.Printf("Got doc with key '%s' from query\n", meta.ID)
		//fmt.Printf("Got doc with key '%s' from query\n", doc)
		result = append(result, doc.Symbol)

	}
	fmt.Println(result)

	//#################################
	//prod := "ProductsList2"
	//ID_ := strconv.Itoa(100)
	ctx = context.Background()
	//getProductbyID := fmt.Sprintf(`FOR product IN ProductsList2 FILTER product.ID == 100_ RETURN product`)
	query2 := "FOR d IN ProductsList2 LIMIT 10 RETURN d"
	//query2 := "FOR product IN ProductsList2 FILTER product.ID == 100_ RETURN product"
	cursor, err = db.Query(ctx, query2, nil)
	if err != nil {
		// handle error
	}
	defer cursor.Close()
	for {
		var doc2 MyDocument

		meta2, err := cursor.ReadDocument(ctx, &doc2)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		if meta2.Key == strconv.Itoa(100){
			fmt.Println(doc2)
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
		//fmt.Println(reflect.TypeOf(meta2.Key))
	}

	//#############################################
	//ID_ := strconv.Itoa(100)
	Symbol := "Symbol"
	ctx = context.Background()
	getProductbyID := fmt.Sprintf(`FOR product IN ProductsList2 FILTER product.%s == @symbol RETURN product`, Symbol)
	//query3 := "FOR product IN ProductsList2 FILTER product.Symbol == @symbol RETURN product"
	bindVars := map[string]interface{}{
		"symbol": "BTC/USD",
	}
	cursor, err = db.Query(ctx, getProductbyID, bindVars)
	if err != nil {
		// handle error
	}
	defer cursor.Close()
	for {
		var doc4 MyDocument

		_, err := cursor.ReadDocument(ctx, &doc4)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
		fmt.Println(doc4)
	}

	//#############################################
	//ID_ := strconv.Itoa(100)
	ID := "ID"
	key := strconv.Itoa(100)
	ctx = context.Background()
	getProductbyID2 := fmt.Sprintf(`FOR product IN ProductsList FILTER product.%s == @id RETURN product`, ID)
	//query3 := "FOR product IN ProductsList2 FILTER product.Symbol == @symbol RETURN product"
	bindVars = map[string]interface{}{
		"id": key,
	}
	cursor, err = db.Query(ctx, getProductbyID2, bindVars)
	if err != nil {
		// handle error
	}
	defer cursor.Close()
	for {
		var produ Product

		_, err := cursor.ReadDocument(ctx, &produ)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		//fmt.Printf("Got doc with key '%s' from query\n", meta2.Key)
		fmt.Println(produ)
	}

}