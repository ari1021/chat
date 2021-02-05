package controller

import (
	"errors"
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/request"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoomHandler struct {
	IRoom model.IRoom
}

func NewRoomHandler(IRoom model.IRoom) RoomHandler {
	return RoomHandler{
		IRoom: IRoom,
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
	ir := rh.IRoom
	r, err := ir.Create(req.Name)
	if err != nil {
		statusCode, res := model.NewAPIResponse(err)
		return c.JSON(statusCode, res)
	}
	// Hubを作成
	h := websocket.NewHub()
	go h.Run()
	model.RoomToHub[r.Model.ID] = h
	return c.JSON(http.StatusOK, r)
}

func (rh RoomHandler) GetRooms(c echo.Context) error {
	ir := rh.IRoom
	rooms, err := ir.FindAll()
	if err != nil {
		res := model.NewAPIError(500, "database error")
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, rooms)
}

func (rh RoomHandler) DeleteRoom(c echo.Context) error {
	// frontからデータを取得
	req := &request.DeleteRoom{}
	if err := c.Bind(req); err != nil {
		return err
	}
	ir := rh.IRoom
	// dbから削除
	r, err := ir.Delete(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := model.NewAPIError(400, "record not found")
			return c.JSON(http.StatusBadRequest, res)
		} else {
			res := model.NewAPIError(500, "database error")
			return c.JSON(http.StatusInternalServerError, res)
		}
	}
	// Hubをstopする
	model.RoomToHub[r.ID].Stop()
	// Hubを削除
	delete(model.RoomToHub, r.ID)
	return c.JSON(http.StatusOK, r)
}
