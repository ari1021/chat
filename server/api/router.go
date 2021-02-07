package api

import (
	"log"

	"github.com/ari1021/websocket/controller"
	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/db"
	"github.com/ari1021/websocket/server/validation"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho(hub *websocket.Hub) *echo.Echo {
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.DB.Conn = conn

	e := echo.New()
	e = validation.ValidateEcho(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "./view/rooms.html")
	e.File("/rooms/create", "./view/create_room.html")
	e.File("/chat", "./view/chat.html")

	e.GET("/ws/:id", controller.ServeRoomWs)
	rr := model.NewRoomRepository(conn)
	rh := controller.NewRoomHandler(rr)
	e.GET("/rooms", rh.GetRooms)
	e.POST("/rooms", rh.CreateRoom)
	e.DELETE("/rooms/:id", rh.DeleteRoom)
	cr := model.NewChatRepository(conn)
	ch := controller.NewChatHandler(cr)
	e.GET("/rooms/:id/chats", ch.GetChats)
	e.POST("/rooms/:id/chats", ch.CreateChat)
	return e
}
