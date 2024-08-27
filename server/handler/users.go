package handler

import (
	"encoding/json"
	"errors"

	// "fmt"
	"net/http"
	"server/api"
	"server/models"
	"strconv"
)

type User struct {
}

var Users = map[string]models.User{}

func (user *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0:
		GetUser(w, r)
		return
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0:
		GetUsers(w, r)
		return
	case r.Method == http.MethodPost:
		CreateUser(w, r)

	}
	// fmt.Fprintf(w, "hello")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if user, exists := Users[id]; exists {
		api.SetResult(http.StatusOK, user, nil, w)
	} else {
		api.SetResult(http.StatusNotFound, nil, nil, w)
	}
	// fmt.Fprintf(w, "user of %s", id)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "users")
	api.SetResult(http.StatusOK, Users, nil, w)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x")
	// for k, v range := r.Header{
	// 	fmt.Println(k, v)
	// }
	if apiKey != "00" {
		// w.WriteHeader(401)
		api.SetResult(http.StatusUnauthorized, nil, nil, w)
		// fmt.Fprintf(w, "invalid api key")
		return
	} else {
		var user *models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			// w.WriteHeader(500)
			// fmt.Fprintf(w, " decode error")
			api.SetResult(http.StatusInternalServerError, nil, err, w)
			return

		} else {
			// w.WriteHeader(200)
			// fmt.Fprintf(w, " user created")
			if _, exists := Users[strconv.Itoa(user.Id)]; exists {
				api.SetResult(http.StatusInternalServerError, nil, errors.New("already exists!"), w)
			} else {
				Users[strconv.Itoa(user.Id)] = *user
				api.SetResult(http.StatusOK, *user, nil, w)
				return

			}
		}

	}

}
