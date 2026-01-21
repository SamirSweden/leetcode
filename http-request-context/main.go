// http request with context in Go 


package main

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"time"
)

func main(){
	
	fmt.Println()


	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	ctx := context.Background()

	ctx , cancel := context.WithTimeout(ctx , time.Second/2)
	defer cancel()

	req , err := http.NewRequestWithContext(ctx , http.MethodGet, "https://google.com", nil)
	if err != nil{
		log.Fatal(err)  // nil context or invalid method
	}


	resp , err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err) // actual request error
	}

	defer resp.Body.Close()


	fmt.Println("status code ", resp.StatusCode)


}






