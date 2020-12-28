package controller

import (
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
)

var seq int = 1

func CreateRoom(c echo.Context) error {
	r := &model.Room{
		ID: seq,
	}
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		r := &model.APIError{
			StatusCode: 400,
			Message:    "room unprocessable entity",
		}
		return c.JSON(http.StatusBadRequest, r)
	}
	model.Rooms[r.ID] = r
	h := websocket.NewHub()
	go h.Run()
	model.RoomToHub[r.ID] = h
	seq += 1
	return c.JSON(http.StatusOK, r)
}

func GetRooms(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Rooms)
}

// func JoinRoom(c echo.Context) error {
// 	roomIDStr := c.Param("id")
// 	roomID, err := strconv.Atoi(roomIDStr)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	hub := model.RoomToHub[roomID]
// 	client := &websocket.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)} //clientを作成して
// 	client.Hub.Register <- client
// 	go client.ReadPump()
// 	go client.WritePump()
// 	return c.JSON(http.StatusOK, roomID)
// }
