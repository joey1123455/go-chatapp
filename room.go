package main

import (
	"log"
	"net/http"

	"github.com/joey1123455/go-chatapp/trace"
	"github.com/stretchr/objx"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

type room struct {
	//forward is a channel that holds incomming messages
	//that should be forwarded to other clients
	forward chan *message
	//join is a channel for clients wishing to join the room
	join chan *client
	//leave is a channel for leaving the room
	leave chan *client
	//clients holds all clients in the room
	clients map[*client]bool
	//avatar is how avater info will be obtained
	avatar Avater
	//tracer recieves information of activity
	tracer trace.Tracer
}

// newRoom makes a new room
func newRoom(avatar Avater) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		avatar:  avatar,
		tracer:  trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			//forward message to all clients
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace("Message recieved: ", msg.Message)
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("serveHTTP:", err)
		return
	}
	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Failed to get auth token")
		return
	}
	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
