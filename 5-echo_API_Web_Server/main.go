package main

import (
	"GoWebFull/4-Building_Optimised_API_Web_Server_UnitTest/cache"
	user "GoWebFull/5-echo_API_Web_Server/users"
	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type jsonResponse map[string]interface{}

func usersOptions(c echo.Context) error{
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	return c.NoContent(http.StatusOK)
}

func userOptions(c echo.Context) error{
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete,
		http.MethodHead, http.MethodOptions}
	c.Response().Header().Set("Allow", strings.Join(methods, ","))
	return c.NoContent(http.StatusOK)
}


func usersGetAll(c echo.Context) error{
	if cache.Serve(c.Response(), c.Request()){
		return nil
	}
	users, err := user.All()
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if c.Request().Method == http.MethodHead{
		return c.NoContent(http.StatusOK)
	}
	c.Response().Writer = cache.NewWriter(c.Response().Writer, c.Request())
	return c.JSON(http.StatusOK, jsonResponse{"users": users})

}

func usersPostOne(c echo.Context) error{
	u := new(user.User)
	err := c.Bind( u)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil{
		if err == user.ErrRecordInvalid{
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	c.Response().Header().Set("Location", "/users/" + u.ID.Hex())
	return c.NoContent(http.StatusCreated)
}

func usersGetOne(c echo.Context) error{
	if cache.Serve(c.Response(), c.Request()){
		return nil
	}
	if !bson.IsObjectIdHex(c.Param("id")){
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	u, err := user.One(id)
	if err != nil{
		if err == storm.ErrNotFound{
			// StatusNotFound = 404 // RFC 7231, 6.5.4
			return echo.NewHTTPError(http.StatusNotFound)
		}
		// StatusInternalServerError = 500 // RFC 7231, 6.6.1
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if c.Request().Method == http.MethodHead{
		return c.NoContent(http.StatusOK)
	}
	// StatusOK = 200 // RFC 7231, 6.3.1
	c.Response().Writer = cache.NewWriter(c.Response().Writer, c.Request())
	return c.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(c echo.Context) error{
	u := new(user.User)
	err := c.Bind( u)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if !bson.IsObjectIdHex(c.Param("id")){
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	u.ID = id
	err = u.Save()
	if err != nil{
		if err == user.ErrRecordInvalid{
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	c.Response().Writer = cache.NewWriter(c.Response().Writer, c.Request())
	return c.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersPatchOne(c echo.Context) error{
	if !bson.IsObjectIdHex(c.Param("id")){
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	u, err := user.One(id)
	if err != nil{
		if err == storm.ErrNotFound{
			// StatusNotFound = 404 // RFC 7231, 6.5.4
			return echo.NewHTTPError(http.StatusNotFound)
		}
		// StatusInternalServerError = 500 // RFC 7231, 6.6.1
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = c.Bind( u)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest)

	}
	u.ID = id
	err = u.Save()
	if err != nil{
		if err == user.ErrRecordInvalid{
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	c.Response().Writer = cache.NewWriter(c.Response().Writer, c.Request())
	return c.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersDeleteOne(c echo.Context) error{
	if !bson.IsObjectIdHex(c.Param("id")){
		return echo.NewHTTPError(http.StatusNotFound)
	}
	id := bson.ObjectIdHex(c.Param("id"))
	err := user.Delete(id)
	if err != nil{
		if err == storm.ErrNotFound{
			// StatusNotFound = 404 // RFC 7231, 6.5.4
			return echo.NewHTTPError(http.StatusNotFound)
		}
		// StatusInternalServerError = 500 // RFC 7231, 6.6.1
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	cache.Drop(cache.MakeResource(c.Request()))
	return c.NoContent(http.StatusOK)
}


func root(c echo.Context) error{
	return c.String(http.StatusOK, "Running API v1")
}

func main(){
	e := echo.New()

	e.GET("/", root)
	u := e.Group("/users")
	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll)
	u.GET("", usersGetAll)
	u.POST("", usersPostOne)

	uid := u.Group("/:id")
	uid.OPTIONS("", userOptions)
	uid.HEAD("", usersGetOne)
	uid.GET("", usersGetOne)
	uid.PUT("", usersPutOne)
	uid.PATCH("", usersPatchOne)
	uid.DELETE("", usersDeleteOne)

	e.Start(":12345")
}

