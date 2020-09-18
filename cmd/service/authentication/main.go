package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/xiashura/server-jwt-example/internal/app/middleware"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func myRegistration(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func myUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<h1>Registration</h1>")
}
func MyDeleteRefresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>DeleteRefresh</h1>")
}

func main() {

	http.HandleFunc("/api/v1/Registration", middleware.Registration(myRegistration))
	http.HandleFunc("/api/v1/Authentication", middleware.Authentication(myHandler))

	http.HandleFunc("/api/v2/Updatetoken", middleware.UpDateToken(myUpdate))

	http.HandleFunc("/api/v3/Refresh/delete/one", middleware.DeleteRefreshToken(myHandler))
	http.HandleFunc("/api/v3/Refresh/delete/many", middleware.DeleteRefreshTokenUser(myHandler))

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
