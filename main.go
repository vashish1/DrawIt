package main

import (
	"DrawIt/hub"
	"DrawIt/user"
	"log"
	"net/http"
)

// type m map[string]*hub.Hub


func main() {
	
var users = []user.User{
	{
		Name: "yashi",
		Slot: hub.NewHub(),
	},
	{
		Name: "Anni",
		Slot: hub.NewHub(),
	},
}

	// var users=&user.Users
	// for _, u := range users {
	// 	u.Slot = hub.NewHub()
	// 	go u.Slot.Run()
		 
	// }
	h1:=users[0].Slot
	h2:=users[1].Slot
	// h := hub.NewHub()
	go h1.Run()
	go h2.Run()
	
	http.HandleFunc("/",h1.HandleWebSocket)
	http.HandleFunc("/connect",h2.HandleWebSocket)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
