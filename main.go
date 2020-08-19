package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	DoRequest()
}

func DoRequest() {
	start := time.Now()
	result, _ := http.Get("https://s3f4.com/")
	fmt.Printf("Request Time: %v, Request Code: %v \n", time.Since(start), result.StatusCode)
}
