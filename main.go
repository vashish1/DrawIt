package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vashish1/DrawIt/hub"
	"github.com/vashish1/DrawIt/user"
)

// type m map[string]*hub.Hub

var p = os.Getenv("PORT")

func main() {

	var users = []user.User{
		{
			Name: "Anurag",
			Slot: hub.NewHub(),
		},
		{
			Name: "Pragati",
			Slot: hub.NewHub(),
		},
		{
			Name: "Tanvi",
			Slot: hub.NewHub(),
		},
		{
			Name: "Priyanka",
			Slot: hub.NewHub(),
		},
		{
			Name: "Dwivedi",
			Slot: hub.NewHub(),
		},
		{
			Name: "Kaushik",
			Slot: hub.NewHub(),
		},
		{
			Name: "Raja",
			Slot: hub.NewHub(),
		},
		{
			Name: "shailendra",
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
	h3 := users[2].Slot
	h4 := users[3].Slot
	h5 := users[4].Slot
	h6 := users[5].Slot
	h7 := users[6].Slot
	h8 := users[7].Slot
	// h := hub.NewHub()
	go h1.Run()
	go h2.Run()

	http.HandleFunc("/index", index)
	http.HandleFunc("/anurag", h1.HandleWebSocket)
	http.HandleFunc("/pragati", h2.HandleWebSocket)
	http.HandleFunc("/tani", h3.HandleWebSocket)
	http.HandleFunc("/priyanka", h4.HandleWebSocket)
	http.HandleFunc("/dwivedi", h5.HandleWebSocket)
	http.HandleFunc("/kaushik", h6.HandleWebSocket)
	http.HandleFunc("/raja", h7.HandleWebSocket)
	http.HandleFunc("/shailendra", h8.HandleWebSocket)

	err := http.ListenAndServe(":"+p, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Working`))
}
