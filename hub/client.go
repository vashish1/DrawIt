package hub

import (
	"github.com/gorilla/websocket"
	uuid "github.com/google/uuid"
	"DrawIt/utility"
)

//Client stores the data w.r.t a client in a hub.
type Client struct {
	id       string
	hub      *Hub
	color    string
	socket   *websocket.Conn
	outbound chan []byte
}

func newClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		id:       uuid.New().String(),
		color:    utility.GenerateColor(),
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (client *Client) read() {
	defer func() {
		client.hub.unregister <- client
	}()
	for {
		_, data, err := client.socket.ReadMessage()
		if err != nil {
			break
		}
		client.hub.onMessage(data, client)
	}
}

func (client *Client) write() {
	for {
		select {
		case data, ok := <-client.outbound:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (client Client) run() {
	go client.read()
	go client.write()
}

func (client Client) close() {
	client.socket.Close()
	close(client.outbound)
}
