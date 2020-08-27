package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xiashura/server-jwt-example/pkg/jwt"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var a jwt.Client

		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {

			fmt.Println("work")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = a.Valid()
		if err != nil {
			fmt.Println("work")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next(w, r)
	}

}

func Mymiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a jwt.Client

		err := json.NewDecoder(r.Body).Decode(&a)

		if err != nil {
			fmt.Println("work")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token := jwt.Registration(a)

		res, _ := json.Marshal(token)

		w.Write(res)
		next(w, r)
	}
}
