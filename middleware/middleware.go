package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/xiashura/server-jwt-example/pkg/jwt"
)

//Registration jwt token
func Registration(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user jwt.Client

		err := json.NewDecoder(r.Body).Decode(&user.User)
		user.User.Time = time.Now()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(user.User.Time.Unix())

		token, err := user.Generate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.Token.Key = token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user.Token)

	}
}

//Authentication jwt token
func Authentication(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var a jwt.Client

		err := json.NewDecoder(r.Body).Decode(&a.Token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = a.Valid()
		fmt.Println(err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next(w, r)
	}

}
