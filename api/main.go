package main

import (
	"net/http"
	"new/test/project/api/metrics"
	"new/test/project/api/router"
)

func main() {
	go metrics.PersistCPUPercentages()
	go metrics.PersistMemoryUsages()
	route := router.New()
	http.ListenAndServe(":8080", route)
}
