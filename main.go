package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/vashish1/DrawIt/hub"
	"github.com/vashish1/DrawIt/user"
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
	h1 := users[0].Slot
	h2 := users[1].Slot
	// h := hub.NewHub()
	go h1.Run()
	go h2.Run()

	http.HandleFunc("/index", index)
	http.HandleFunc("/", h1.HandleWebSocket)
	http.HandleFunc("/connect", h2.HandleWebSocket)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./Index.html")
	if err != nil {
		log.Fatal("Could not parse template files\n", err)
	}
	er := t.Execute(w, "")
	if er != nil {
		log.Fatal("could not execute the files\n", er)
	}
}
