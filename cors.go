

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


type CheckRequest struct{
	Host string `json:"host"`
}


type CheckResponse struct {
	Host     string `json:"host"`
	Status   string `json:"status"`
	Duration string `json:"duration"`
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/check" , CheckHostHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080" , CorsPolicy(mux))
}



func CheckHostHandler(w http.ResponseWriter ,r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, " method not allowed" , http.StatusMethodNotAllowed)
		return 
	}

	var req CheckRequest 
	json.NewDecoder(r.Body).Decode(&req)

	if req.Host == "" {
		http.Error(w, "url is required" , http.StatusBadRequest)
		return
	}

	start := time.Now()
	client := http.Client {
		Timeout: 3 * time.Second,
	}
	resp , err := 	client.Get("https://" + req.Host)
	duration := time.Since(start)

	result := CheckResponse{Host : req.Host , Duration: duration.String()}

	if err != nil {
		result.Status = "❌ host is down"
	} else {
		result.Status = fmt.Sprintf("✅ Website is alive (HTTP %d)" , resp.StatusCode)
		resp.Body.Close()
	}

	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(result)
}

func CorsPolicy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin" , "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods" , "GET , POST")
		w.Header().Set("Access-Control-Allow-Headers" ,  "Content-Type")


		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w,r)
	})
}
