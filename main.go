package main

import (
	"log"
  "net/http"
  "DrawIt/hub"
)

func main() {
	h := hub.NewHub()
	go h.Run()
	http.HandleFunc("/ws", h.HandleWebSocket)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
