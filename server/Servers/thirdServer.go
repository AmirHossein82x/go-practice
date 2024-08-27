package Servers

import (
	// "context"
	"fmt"
	"net/http"
	"time"
)

func ThirdServer() {
	mx := http.NewServeMux()
	mx.HandleFunc("/users", HandleUsers)
	server := http.Server{
		Addr:    ":8080",
		Handler: mx,
	}
	server.ListenAndServe()
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("default"))
	ctx := r.Context()
	time.Sleep(time.Second * 5)
	select {
	case <-ctx.Done():
		w.Write([]byte("canceled by user"))
		fmt.Println("canceled by user")
	case <-time.After(3 * time.Second):
		w.Write([]byte("response"))
	default:
		w.Write([]byte("default"))
	}
}
