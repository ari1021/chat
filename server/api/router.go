package api

import (
	"log"
	"strconv"

	"github.com/ari1021/websocket/controller"
	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func NewEcho(hub *websocket.Hub) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "./view/rooms.html")
	e.File("/rooms/create", "./view/create_room.html")
	e.File("/chat", "./view/chat.html")

	e.GET("/ws/:id", func(c echo.Context) error {
		// pathparamのgroupIdを取得
		// groupID->*hubを取得
		roomIDStr := c.Param("id")
		roomID, err := strconv.Atoi(roomIDStr)
		if err != nil {
			log.Println(err)
			return err
		}
		hub := model.RoomToHub[roomID]
		controller.ServeWs(hub, c)
		return nil
	})
	e.Validator = &customValidator{Validator: validator.New()}
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.CreateUser)
	e.GET("/rooms", controller.GetRooms)
	e.POST("/rooms", controller.CreateRoom)
	return e
}
