package main

import (
	"fmt"
	"net/http"
)

func main(){
	res , err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}


	fmt.Println("Status code" , res.StatusCode)

}




