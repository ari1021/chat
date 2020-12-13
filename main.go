// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"os"

	"github.com/ari1021/websocket/controller"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.File("/", "./view/index.html")
	e.GET("/ws", func(c echo.Context) error {
		controller.ServeWs(hub, c)
		return nil
	})
	e.Validator = &customValidator{Validator: validator.New()}
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.CreateUser)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
