package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vashish1/DrawIt/hub"
	"github.com/vashish1/DrawIt/user"
)

// type m map[string]*hub.Hub

var Port = os.Getenv("PORT")

func main() {

	var r = mux.NewRouter()
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
	go h3.Run()
	go h4.Run()
	go h5.Run()
	go h6.Run()
	go h7.Run()
	go h8.Run()

	r.HandleFunc("/index", index)
	r.HandleFunc("/anurag", h1.Anurag)
	r.HandleFunc("/pragati", h2.Pragati)
	r.HandleFunc("/tani", h3.Tanvi)
	r.HandleFunc("/priyanka", h4.Priyanka)
	r.HandleFunc("/dwivedi", h5.Dwivedi)
	r.HandleFunc("/kaushik", h6.Kaushik)
	r.HandleFunc("/raja", h7.Raja)
	r.HandleFunc("/shailendra", h8.Shailendra)

	http.Handle("/", r)
	http.ListenAndServe(":"+Port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Working`))
}