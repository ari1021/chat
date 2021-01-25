package controller

import (
	"net/http"

	"github.com/ari1021/websocket/db"
	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/request"
	"github.com/labstack/echo/v4"
)

func CreateChat(c echo.Context) error {
	req := &request.CreateChat{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		res := model.NewAPIError(400, "chat unprocessable entity")
		return c.JSON(http.StatusBadRequest, res)
	}
	conn := db.DB.GetConnection()
	chat := &model.Chat{
		RoomID:  req.RoomID,
		UserID:  req.UserID,
		Message: req.Message,
	}
	if _, err := chat.Create(conn); err != nil {
		statusCode, res := model.NewAPIResponse(err)
		return c.JSON(statusCode, res)
	}
	return c.JSON(http.StatusOK, chat)
}

func GetChats(c echo.Context) error {
	req := &request.GetChats{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		res := model.NewAPIError(400, "limit or offset not found")
		return c.JSON(http.StatusBadRequest, res)
	}
	conn := db.DB.GetConnection()
	chats := &model.Chats{}
	chats, err := chats.Find(conn, req.RoomID, req.Limit, req.Offset)
	if err != nil {
		res := model.NewAPIError(500, "database error")
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, chats)
}
