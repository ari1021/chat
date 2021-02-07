package controller

import (
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/request"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	IChat model.IChat
}

func NewChatHandler(IChat model.IChat) ChatHandler {
	return ChatHandler{
		IChat: IChat,
	}
}

func (ch ChatHandler) CreateChat(c echo.Context) error {
	req := &request.CreateChat{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		res := model.NewAPIError(400, "chat unprocessable entity")
		return c.JSON(http.StatusBadRequest, res)
	}
	ic := ch.IChat
	chat, err := ic.Create(req.Message, req.RoomID, req.UserName)
	if err != nil {
		statusCode, res := model.NewAPIResponse(err)
		return c.JSON(statusCode, res)
	}
	return c.JSON(http.StatusOK, chat)
}

func (ch ChatHandler) GetChats(c echo.Context) error {
	req := &request.GetChats{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		res := model.NewAPIError(400, "limit not found")
		return c.JSON(http.StatusBadRequest, res)
	}
	ic := ch.IChat
	chats, err := ic.Find(req.RoomID, req.Limit, req.Offset)
	if err != nil {
		res := model.NewAPIError(500, "database error")
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, chats)
}
