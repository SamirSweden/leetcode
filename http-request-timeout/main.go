package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func main(){
	
	httpClient := &http.Client{
		Transport: nil,
		CheckRedirect: nil,
		Jar: nil,

		Timeout: 2 * time.Second,
	}


	res , err := httpClient.Get("https://google.com/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status code [ " ,res.StatusCode, " ]")

}

// 200 ok 


