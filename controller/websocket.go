package controller

import (
	"log"

	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
)

// serveWs handles websocket requests from the peer.
func ServeWs(hub *websocket.Hub, c echo.Context) {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &websocket.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)} //clietを作成して
	client.Hub.Register <- client                                                   //Hubにregisterする

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
