package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on Connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {

			log.Println(msg)

			log.Println("emit:", so.Emit("chat message", msg))

			so.BroadcastTo("chat", "chat message", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		server.ServeHTTP(w, r)
	})

	//http.Handle("/socket.io/", server)

	http.Handle("/", http.FileServer(http.Dir("./asset")))

	log.Println("Serving at localhost:5000...")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
