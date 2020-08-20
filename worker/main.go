package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("test")
	i := 0
	for i < 1 {
		fmt.Println("test")
		DoRequest()
		i++
	}
}

func DoRequest() {
	start := time.Now()
	result, err := http.Get("https://s3f4.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Request Time: %v, Request Code: %v \n", time.Since(start), result.StatusCode)
}
