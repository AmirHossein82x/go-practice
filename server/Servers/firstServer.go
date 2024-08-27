package Servers

import (
	"net/http"
	"time"
)


func FirstServer() {
	server := http.Server{
		Addr: ":8080",
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 5,
	}
	server.ListenAndServe()
}