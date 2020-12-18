package controller

import (
	"net/http"

	"github.com/ari1021/websocket/model"
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
	seq += 1
	return c.JSON(http.StatusOK, r)
}

func GetRooms(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Rooms)
}
