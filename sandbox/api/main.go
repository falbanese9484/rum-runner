package main

import (
	"fmt"
	"net/http"

	"github.com/falbanese9484/rum-runner"
)

const (
	API_KEY = "someapi"
)

func MiddlewareTest(c *rum.RumContext) {
	apikey := c.R.Header.Get("x-api-key")
	if apikey != API_KEY {
		c.JSON(401, map[string]any{
			"error": "unauthorized",
		})
		return
	} else {
		c.Next()
	}
}

func TestingSomething(c *rum.RumContext) {
	c.JSON(200, map[string]any{
		"request-id": c.RequestId(),
		"success":    "ok",
	})
}

func main() {
	e := rum.New()
	e.Use(MiddlewareTest)
	v1 := e.NewGroup("/api/v1")
	v1.GET("/", TestingSomething)

	port := ":8945"

	fmt.Printf("Listening on Port: %s\n", port)
	http.ListenAndServe(port, e)
}
