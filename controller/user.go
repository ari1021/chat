package controller

import (
	"net/http"

	"github.com/ari1021/websocket/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func createUser(c echo.Context) error {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u := &model.User{
		ID: uuidObj.String(), //uuidで生成
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	model.Users[u.ID] = u
	return c.JSON(http.StatusOK, u)
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Users)
}
