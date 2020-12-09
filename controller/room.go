package controller

import (
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/labstack/echo/v4"
)

var seq int = 1

func createRoom(c echo.Context) error {
	g := &model.Group{
		ID: seq,
	}
	if err := c.Bind(g); err != nil {
		return err
	}
	if err := c.Validate(g); err != nil {
		return err
	}
	model.Groups[g.ID] = g
	seq += 1
	return c.JSON(http.StatusOK, g)
}
