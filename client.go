package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	//socket is websocket for this client
	socket *websocket.Conn
	//send is a channel on which messages are sent
	send chan *message
	//room is the room client is chatting in
	room *room
	//user data holds informatin about the client
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		msg.AvatarURL, _ = c.room.avatar.GetAvaterURL(c)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
