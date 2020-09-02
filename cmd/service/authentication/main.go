package main

import (
	"fmt"
	"net/http"

	"github.com/xiashura/server-jwt-example/middleware"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Home</h1>")
}

func myRegistration(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Registration</h1>")
}

func main() {
	http.HandleFunc("/", middleware.Authentication(myHandler))
	http.HandleFunc("/Registration", middleware.Registration(myRegistration))
	http.HandleFunc("/Authentication", middleware.Authentication(myHandler))
	http.HandleFunc("/Unauthenticated", middleware.Authentication(myHandler))
	http.HandleFunc("/Expired", middleware.Authentication(myHandler))
	http.ListenAndServe(":8080", nil)
}
