package controller

import (
	"log"
	"strconv"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
)

func ServeRoomWs(c echo.Context) error {
	// pathparamのgroupIdを取得
	// groupID->*hubを取得
	roomIDStr := c.Param("id")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		return err
	}
	hub := model.RoomToHub[uint(roomID)]
	serveWs(hub, c)
	return nil
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *websocket.Hub, c echo.Context) {
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
