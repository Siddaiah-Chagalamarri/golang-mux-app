package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type RequestBody struct {
	Text string `json:"text"`
}

type ResponseBody struct {
	Result string `json:"result"`
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	hashedText, err := bcrypt.GenerateFromPassword([]byte(reqBody.Text), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	resp := ResponseBody{
		Result: string(hashedText),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resp := ResponseBody{
		Result: reqBody.Text,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")
	r.HandleFunc("/api/decrypt", DecryptHandler).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
