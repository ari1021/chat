package controller

import (
	"errors"
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/db"
	"github.com/ari1021/websocket/server/request"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoomHandler struct {
	db   db.DBConnection
	room *model.Room
}

func NewRoomHandler(db db.DBConnection, room *model.Room) RoomHandler {
	return RoomHandler{
		db:   db,
		room: room,
	}
}

func (rh RoomHandler) CreateRoom(c echo.Context) error {
	// frontからデータを取得
	req := &request.CreateRoom{}
	if err := c.Bind(req); err != nil {
		return err
	}
	// validationを行う
	if err := c.Validate(req); err != nil {
		res := model.NewAPIError(400, "room unprocessable entity")
		return c.JSON(http.StatusBadRequest, res)
	}
	// dbに保存
	conn := rh.db.GetConnection()
	r := rh.room
	if _, err := r.Create(conn); err != nil {
		statusCode, res := model.NewAPIResponse(err)
		return c.JSON(statusCode, res)
	}
	// Hubを作成
	h := websocket.NewHub()
	go h.Run()
	model.RoomToHub[r.Model.ID] = h
	return c.JSON(http.StatusOK, r)
}

// func CreateRoom(c echo.Context) error {
// 	// frontからデータを取得
// 	req := &request.CreateRoom{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}
// 	// validationを行う
// 	if err := c.Validate(req); err != nil {
// 		res := model.NewAPIError(400, "room unprocessable entity")
// 		return c.JSON(http.StatusBadRequest, res)
// 	}
// 	// dbに保存
// 	conn := db.DB.GetConnection()
// 	r := &model.Room{
// 		Name: req.Name,
// 	}
// 	if _, err := r.Create(conn); err != nil {
// 		statusCode, res := model.NewAPIResponse(err)
// 		return c.JSON(statusCode, res)
// 	}
// 	// Hubを作成
// 	h := websocket.NewHub()
// 	go h.Run()
// 	model.RoomToHub[r.Model.ID] = h
// 	return c.JSON(http.StatusOK, r)
// }

func GetRooms(c echo.Context) error {
	r := &model.Rooms{}
	conn := db.DB.GetConnection()
	rooms, err := r.FindAll(conn)
	if err != nil {
		res := model.NewAPIError(500, "database error")
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
			res := model.NewAPIError(400, "record not found")
			return c.JSON(http.StatusBadRequest, res)
		} else {
			res := model.NewAPIError(500, "database error")
			return c.JSON(http.StatusInternalServerError, res)
		}
	}
	// Hubをstopする
	model.RoomToHub[r.Model.ID].Stop()
	// Hubを削除
	delete(model.RoomToHub, r.Model.ID)
	return c.JSON(http.StatusOK, r)
}
