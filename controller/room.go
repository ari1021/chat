package controller

import (
	"log"
	"net/http"

	"github.com/ari1021/websocket/db"
	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/request"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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
	conn := db.DB.GetConnection()
	r := &model.Room{
		Name:   req.Name,
		UserID: req.UserID,
	}
	if _, err := r.Create(conn); err != nil {
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
	conn := db.DB.GetConnection()
	rooms, err := r.GetAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, rooms)
}

func DeleteRoom(c echo.Context) error {
	// frontからデータを取得
	req := &request.DeleteRoom{}
	if err := c.Bind(req); err != nil {
		return err
	}
	// validationを行う
	if err := c.Validate(req); err != nil {
		r := &model.APIError{
			StatusCode: 400,
			Message:    "room_id bad entity",
		}
		return c.JSON(http.StatusBadRequest, r)
	}
	conn := db.DB.GetConnection()
	r := &model.Room{
		Model: gorm.Model{
			ID: req.ID,
		},
	}
	// dbから削除
	if _, err := r.Delete(conn); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, r)
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
