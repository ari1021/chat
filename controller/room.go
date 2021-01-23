package controller

import (
	"errors"
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
		res := &model.APIError{
			StatusCode: 400,
			Message:    "room unprocessable entity",
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// dbに保存
	conn := db.DB.GetConnection()
	r := &model.Room{
		Name:   req.Name,
		UserID: req.UserID,
	}
	if _, err := r.Create(conn); err != nil {
		if model.IsForeignKeyError(err) {
			res := &model.APIError{
				StatusCode: 400,
				Message:    "foreign key error",
			}
			return c.JSON(http.StatusBadRequest, res)
		} else if model.IsDuplicateKeyError(err) {
			res := &model.APIError{
				StatusCode: 400,
				Message:    "duplicate key error",
			}
			return c.JSON(http.StatusBadRequest, res)
		} else {
			res := &model.APIError{
				StatusCode: 500,
				Message:    "database error",
			}
			return c.JSON(http.StatusInternalServerError, res)
		}
	}
	// Hubを作成
	h := websocket.NewHub()
	go h.Run()
	model.RoomToHub[r.Model.ID] = h
	return c.JSON(http.StatusOK, r)
}

func GetRooms(c echo.Context) error {
	r := &model.Rooms{}
	conn := db.DB.GetConnection()
	rooms, err := r.FindAll(conn)
	if err != nil {
		res := &model.APIError{
			StatusCode: 500,
			Message:    "database error",
		}
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, rooms)
}

func DeleteRoom(c echo.Context) error {
	// frontからデータを取得
	req := &request.DeleteRoom{}
	if err := c.Bind(req); err != nil {
		return err
	}
	conn := db.DB.GetConnection()
	r := &model.Room{
		Model: gorm.Model{
			ID: uint(req.ID),
		},
	}
	// dbから削除
	if _, err := r.Delete(conn); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := &model.APIError{
				StatusCode: 400,
				Message:    "record not found",
			}
			return c.JSON(http.StatusBadRequest, res)
		} else {
			res := &model.APIError{
				StatusCode: 500,
				Message:    "database error",
			}
			return c.JSON(http.StatusInternalServerError, res)
		}
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
