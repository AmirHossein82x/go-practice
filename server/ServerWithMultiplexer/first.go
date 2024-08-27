package ServerWithMultiplexer

import (
	"net/http"
	"time"
)

func CreateFirstServerWithMux() {
	mux := http.NewServeMux()
	mux.Handle("/google", http.RedirectHandler("https://google.com", 307))
	mux.Handle("/digikala", http.RedirectHandler("https://digikala.com", 307))
	mux.HandleFunc("/get-users",GetUsers)
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}

	server.ListenAndServe()

}

func GetUsers(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("Peyman"))
}