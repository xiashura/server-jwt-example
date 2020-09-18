package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xiashura/server-jwt-example/internal/app/db"
	"github.com/xiashura/server-jwt-example/internal/app/mycrypt"
	"github.com/xiashura/server-jwt-example/internal/app/tokens"
	"github.com/xiashura/server-jwt-example/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Registration jwt token
func Registration(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := db.Data{
			User: user,
		}

		data.User.ID = primitive.NewObjectID()
		res, err := data.Create_User_db()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)

	}

}

//Authentication jwt token
func Authentication(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var a tokens.Client
		//var db db.Data
		err := json.NewDecoder(r.Body).Decode(&a.Token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ID := a.ValidHash()

		datadb := db.Data{}

		res, err := datadb.Find_User_db_by_id(ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashToken := base64.StdEncoding.EncodeToString([]byte(mycrypt.Mycrypt(a.Token.Refresh)))

		for _, el := range res.HashTokens {
			if string(el) != hashToken {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		next(w, r)

	}

}

func UpDateToken(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user db.Data

		err := json.NewDecoder(r.Body).Decode(&user.User)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := user.Find_User_db()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		clinet := tokens.Client{
			User: res,
		}

		tokens, err := clinet.Generate(res.ID)
		user.Append_token(tokens.Refresh)

		/// add hash tokens

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(tokens)

	}

}

//DeleteRefreshToken on db
func DeleteRefreshToken(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user db.Data

		err := json.NewDecoder(r.Body).Decode(&user.Token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := user.Delete_Refresh_Token()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(res)

	}

}

//DeleteRefreshTokenUser delete all key in user model
func DeleteRefreshTokenUser(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user db.Data

		err := json.NewDecoder(r.Body).Decode(&user.User)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := user.Delete_AllRefresh_Token()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(res)

	}

}
