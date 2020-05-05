package main

import (
	"github.com/labstack/echo"
	"net/http"
)


func root(c echo.Context) error{
	return c.String(http.StatusOK, "Running API v1")
}

func main(){
	e := echo.New()

	e.GET("/", root)
	e.Start(":12345")
}

