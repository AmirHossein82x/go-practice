package ServerWithMultiplexer

import (
	// "fmt"
	"net/http"
	"time"
)

type TestHandler struct {
}

func CreateFirstServerWithCustomHandler() {

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 5,
		Handler:      TestHandler{},
	}

	server.ListenAndServe()

}

func (testHandler TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "hello world")
	w.WriteHeader(401)
	w.Write([]byte("hello world"))
}
