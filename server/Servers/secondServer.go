package Servers

import "net/http"

func SecondServer() {
	http.ListenAndServe(":8090", nil)
}