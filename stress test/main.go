package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const URL = "http://example.com"

func main() {
	var wg sync.WaitGroup
	
	fmt.Println("🚀 sending HTTP requests...")

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(URL)
			elapsed := time.Since(start)

			if err != nil {
				fmt.Printf("❌ [%2d] Error: %v\n", id, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("✅ [%2d] %d | %v\n", id, resp.StatusCode, elapsed.Round(time.Millisecond))
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Test completed !")
}
