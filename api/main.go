package main

import (
	"net/http"
	"new/test/project/api/router"
)

func main() {
	route := router.New()
	http.ListenAndServe(":8080", route)
}
