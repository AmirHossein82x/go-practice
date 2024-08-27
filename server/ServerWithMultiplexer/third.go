package ServerWithMultiplexer

import (
	"net/http"
	"server/handler"
	"time"
)

func CreateFirstServerWithMux2() {
	mux := http.NewServeMux()

	mux.Handle("/get-users", &handler.User{})
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}

	server.ListenAndServe()

}

