package controller

import (
	"log"
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/request"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
)

var seq int = 1

func CreateRoom(c echo.Context) error {
	// frontからデータを取得
	req := &request.CreateRoom{}
	if err := c.Bind(req); err != nil {
		return err
	}
	// validationを行う
	if err := c.Validate(req); err != nil {
		r := &model.APIError{
			StatusCode: 400,
			Message:    "room unprocessable entity",
		}
		return c.JSON(http.StatusBadRequest, r)
	}
	// dbに保存
	r := &model.Room{
		Name:   req.Name,
		UserID: req.UserId,
	}
	if _, err := r.Create(); err != nil {
		log.Fatal(err)
	}
	// Hubを作成
	h := websocket.NewHub()
	go h.Run()
	model.RoomToHub[r.Model.ID] = h
	return c.JSON(http.StatusOK, r)
}

func GetRooms(c echo.Context) error {
	r := &model.Room{}
	rooms, err := r.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, rooms)
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
