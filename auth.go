package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET , POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func register(w http.ResponseWriter, r *http.Request) {
	setCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	if name == "" || password == "" {
		http.Error(w, "Register please", http.StatusBadRequest)
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}

	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "error while creating file", http.StatusInternalServerError)
		return
	}

	_, err = f.WriteString(fmt.Sprintf("%s:%s\n", name, string(hashedpassword)))
	if err != nil {
		http.Error(w, "cannot write to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "registered successfully")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	setCors(w)
	fmt.Fprintf(w, "Welcome to dashboard")
}

func main() {

	http.HandleFunc("/register", register)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
